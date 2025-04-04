package handlers

import (
	"net/http"
	"strconv"
	"sync"
	"web-api-in-go/models"

	"github.com/labstack/echo/v4"
)

// sync.Map を定義
var taskStore sync.Map
var taskId int

// Task 全取得ハンドラ
func GetTasks(c echo.Context) error {
	var tasks []models.Task

	taskStore.Range(func(key, value interface{}) bool {
		tasks = append(tasks, value.(models.Task))
		return true
	})

	if len(tasks) > 0 {
		return c.JSON(http.StatusOK, tasks)
	}
	return c.String(http.StatusOK, "GetTasks Called!!")
}

// Task 1件取得ハンドラ
func GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // paramから取得したIDをintに変換
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	var matchedTask *models.Task
	taskStore.Range(func(key, value interface{}) bool {
		task := value.(models.Task)
		if task.ID == id {
			matchedTask = &task
			return false
		}
		return true
	})
	if matchedTask != nil {
		return c.JSON(http.StatusOK, matchedTask)
	} else {
		return c.String(http.StatusNotFound, "Task not found")
	}
}

// Task 作成ハンドラ
func CreateTask(c echo.Context) error {
	taskId++
	newTask := models.Task{
		ID:        taskId,
		Title:     c.FormValue("title"),
		Completed: false,
	}

	taskStore.Store(taskId, newTask)
	return c.String(http.StatusOK, "Task Created!!")
}

// Task 更新ハンドラ
func UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // paramから取得したIDをintに変換
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	var matchedTask *models.Task
	taskStore.Range(func(key, value interface{}) bool {
		task := value.(models.Task)
		if task.ID == id {
			matchedTask = &task
			return false
		}
		return true
	})

	if matchedTask != nil {
		title := matchedTask.Title
		if c.FormValue("title") != "" {
			title = c.FormValue("title")
		}
		// taskStore 内の値を更新する
		taskStore.Store(id, models.Task{
			ID:        id,
			Title:     title,
			Completed: c.FormValue("completed") == "true",
		})
		return c.String(http.StatusOK, "Task Updated!!")
	} else {
		return c.String(http.StatusNotFound, "Task not found")
	}
}

// Task 削除ハンドラ
func DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // paramから取得したIDをintに変換
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	taskStore.Delete(id)

	return c.String(http.StatusOK, "Task Deleted!!")
}