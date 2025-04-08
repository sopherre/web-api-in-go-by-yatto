package handlers

import (
	"net/http"
	"strconv"
	"web-api-in-go/db"
	"web-api-in-go/models"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	db *db.Database
}

func NewTaskHandler(db *db.Database) *TaskHandler {
	return &TaskHandler{db: db}
}

func (h *TaskHandler) SetupRoutes(e *echo.Echo) {
	e.GET("/tasks", h.GetTasks)
	e.GET("/tasks/:id", h.GetTask)
	e.POST("/tasks", h.CreateTask)
	e.PUT("/tasks/:id", h.UpdateTask)
	e.DELETE("/tasks/:id", h.DeleteTask)
}

// Task 全取得ハンドラ
func (h *TaskHandler) GetTasks(c echo.Context) error {
	var tasks []models.Task
	result := h.db.GormDb.Find(&tasks)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error getting tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}

// Task 1件取得ハンドラ
func (h *TaskHandler) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // paramから取得したIDをintに変換
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	var task models.Task
	result := h.db.GormDb.First(&task, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}
	return c.JSON(http.StatusOK, task)
}

// Task 作成ハンドラ
func (h *TaskHandler) CreateTask(c echo.Context) error {
	newTask := models.Task{}
	if err := c.Bind(&newTask); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	result := h.db.GormDb.Create(&newTask)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error creating task")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Task Created!!", "id": newTask.ID})
}

// Task 更新ハンドラ
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	// IDで既存のタスクを取得
	var task models.Task
	result := h.db.GormDb.First(&task, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}

	// リクエストボディから更新データを取得
	var input models.Task
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// 値を更新（空チェックなど必要なら追加）
	if input.Title != "" {
		task.Title = input.Title
	}
	task.Completed = input.Completed

	// DB保存
	result = h.db.GormDb.Save(&task)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error updating task")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Task Updated!!",
		"id":      task.ID,
	})
}

// Task 削除ハンドラ
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	var task models.Task
	result := h.db.GormDb.First(&task, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}
	result = h.db.GormDb.Delete(&task)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error deleting task")
	}
	return c.String(http.StatusOK, "Task Deleted!!")
}
