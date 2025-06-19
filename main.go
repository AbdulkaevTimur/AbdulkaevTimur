package main

import (
	"AbdulkaevTimur/task"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var manager = task.NewManager()

func main() {
	e := echo.New()

	e.POST("/tasks", handleCreateTask)
	e.GET("/tasks/:id", handleGetTask)
	e.DELETE("/tasks/:id", handleDeleteTask)

	e.Logger.Fatal(e.Start(":8080"))
}

func handleCreateTask(c echo.Context) error {
	id := manager.CreateTask()
	return c.JSON(http.StatusOK, map[string]string{"task_id": id})
}

func handleGetTask(c echo.Context) error {
	id := c.Param("id")
	t, err := manager.GetTask(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}

	resp := struct {
		ID          string `json:"id"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
		StartedAt   string `json:"started_at,omitempty"`
		FinishedAt  string `json:"finished_at,omitempty"`
		DurationSec int64  `json:"duration_sec"`
		Result      string `json:"result,omitempty"`
	}{
		ID:          t.ID,
		Status:      string(t.Status),
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
		StartedAt:   t.StartedAt.Format(time.RFC3339),
		FinishedAt:  t.FinishedAt.Format(time.RFC3339),
		DurationSec: int64(t.Duration().Seconds()),
		Result:      t.Result,
	}
	return c.JSON(http.StatusOK, resp)
}

func handleDeleteTask(c echo.Context) error {
	id := c.Param("id")
	err := manager.DeleteTask(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
	}
	return c.NoContent(http.StatusNoContent)
}
