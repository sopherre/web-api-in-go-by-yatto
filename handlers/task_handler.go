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

// GetTasks godoc
// @Summary タスク一覧を取得
// @Description 登録されている全てのタスクを取得します
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {string} string "Error getting tasks"
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(c echo.Context) error {
	var tasks []models.Task
	result := h.db.GormDb.Find(&tasks)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error getting tasks")
	}
	return c.JSON(http.StatusOK, tasks)
}

// GetTask godoc
// @Summary タスクを1件取得
// @Description 指定したIDのタスク情報を取得します
// @Tags tasks
// @Produce json
// @Param id path int true "タスクID"
// @Success 200 {object} models.Task
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
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

// CreateTask godoc
// @Summary タスクを作成
// @Description 新しいタスクを作成します
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "作成するタスク情報"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Error creating task"
// @Router /tasks [post]
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

// UpdateTask godoc
// @Summary タスクを更新
// @Description 指定したIDのタスク情報を更新します
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "タスクID"
// @Param task body models.Task true "更新するタスク情報"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Invalid ID or request body"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Error updating task"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	var task models.Task
	result := h.db.GormDb.First(&task, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Task not found")
	}
	var input models.Task
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}
	if input.Title != "" {
		task.Title = input.Title
	}
	task.Completed = input.Completed

	result = h.db.GormDb.Save(&task)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error updating task")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Task Updated!!",
		"id":      task.ID,
	})
}

// DeleteTask godoc
// @Summary タスクを削除
// @Description 指定したIDのタスクを削除します
// @Tags tasks
// @Produce plain
// @Param id path int true "タスクID"
// @Success 200 {string} string "Task Deleted!!"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Task not found"
// @Failure 500 {string} string "Error deleting task"
// @Router /tasks/{id} [delete]
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
