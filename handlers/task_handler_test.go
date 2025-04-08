package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"web-api-in-go/db"
	"web-api-in-go/handlers"
	"web-api-in-go/models"
)

type TaskHandlerTestSuite struct {
	suite.Suite
	echo        *echo.Echo
	taskHandler *handlers.TaskHandler
	db          *db.Database
	task        models.Task
	taskList    []models.Task
}

func (s *TaskHandlerTestSuite) SetupTest() {
	s.echo = echo.New()
	s.db = db.Init()
	// トランザクション開始
	tx := s.db.GormDb.Begin()
	s.db.GormDb = tx // ハンドラと共有する DB をトランザクションに切り替え

	s.taskHandler = handlers.NewTaskHandler(s.db)

	s.task = models.Task{Title: "test task", Completed: false}
	tx.Create(&s.task)

	s.taskList = []models.Task{
		{Title: "test task 1", Completed: false},
		{Title: "test task 2", Completed: false},
	}
	tx.Create(&s.taskList)
}

func (s *TaskHandlerTestSuite) TearDownTest() {
	s.db.GormDb.Rollback()
}

func (s *TaskHandlerTestSuite) TestGetTasks() {
	req, _ := http.NewRequest(echo.GET, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	s.taskHandler.GetTasks(c)
	s.Equal(http.StatusOK, rec.Code)

	var responseTasks []models.Task
	s.NoError(json.Unmarshal(rec.Body.Bytes(), &responseTasks))

	var dbTasks []models.Task
	s.db.GormDb.Order("id asc").Find(&dbTasks)

	s.Equal(dbTasks, responseTasks)
}

func (s *TaskHandlerTestSuite) TestGetTask() {
	req, _ := http.NewRequest(echo.GET, fmt.Sprintf("/tasks/%d", s.task.ID), nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(s.task.ID))

	s.taskHandler.GetTask(c)
	s.Equal(http.StatusOK, rec.Code)

	var responseTask models.Task
	s.NoError(json.Unmarshal(rec.Body.Bytes(), &responseTask))

	var dbTask models.Task
	s.db.GormDb.First(&dbTask, s.task.ID)

	s.Equal(dbTask, responseTask)
}

func (s *TaskHandlerTestSuite) TestCreateTask() {
	newTask := models.Task{
		Title:     "new task",
		Completed: false,
	}
	body, _ := json.Marshal(newTask)

	req, _ := http.NewRequest(echo.POST, "/tasks", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // JSONとして扱うため必要
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)

	err := s.taskHandler.CreateTask(c)
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)

	// レスポンスの内容確認
	var res map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	s.NoError(err)
	s.Equal("Task Created!!", res["message"])
	s.NotNil(res["id"])
	fmt.Println(res["id"])

	// DBからタスクを取得して確認
	var dbTask models.Task
	result := s.db.GormDb.First(&dbTask, res["id"])
	s.NoError(result.Error)
	s.Equal(newTask.Title, dbTask.Title)
	s.Equal(newTask.Completed, dbTask.Completed)
}

func (s *TaskHandlerTestSuite) TestUpdateTask() {
	// Completedだけ更新する入力用タスク（Titleは空でもOK）
	inputTask := models.Task{
		Completed: true,
	}
	body, _ := json.Marshal(inputTask)

	req, _ := http.NewRequest(echo.PUT, fmt.Sprintf("/tasks/%d", s.task.ID), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // ★ 重要！

	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(s.task.ID))

	// 実行
	err := s.taskHandler.UpdateTask(c)
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)

	// 更新後のレコード取得
	var updated models.Task
	result := s.db.GormDb.First(&updated, s.task.ID)
	s.NoError(result.Error)

	// アサーション
	s.Equal(s.task.Title, updated.Title) // タイトルは変わってないこと
	s.Equal(true, updated.Completed)     // Completed が更新されていること
}

func (s *TaskHandlerTestSuite) TestDeleteTask() {
	req, _ := http.NewRequest(echo.DELETE, fmt.Sprintf("/tasks/%d", s.task.ID), nil)
	rec := httptest.NewRecorder()
	c := s.echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(s.task.ID))

	s.taskHandler.DeleteTask(c)
	s.Equal(http.StatusOK, rec.Code)

	var task models.Task
	err := s.db.GormDb.First(&task, s.task.ID).Error
	s.Error(err)
	s.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func TestTaskHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskHandlerTestSuite))
}
