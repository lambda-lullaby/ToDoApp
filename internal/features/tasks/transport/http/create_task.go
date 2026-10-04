package tasks_transport_http

import (
	"net/http"

	"github.com/google/uuid"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type CreateTaskRequest struct {
	Title        string    `json:"title"          validate:"required,min=1,max=100"`
	Description  *string   `json:"description"    validate:"omitempty,min=1,max=1000"`
	AuthorUserID uuid.UUID `json:"author_user_id" validate:"required"`
}

func (h *TasksHTTPHandler) CreateTask(c *core_http.Context) error {
	var req CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	task, err := h.tasksService.CreateTask(c.Context(), req.Title, req.Description, req.AuthorUserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, taskToResponse(task))
}
