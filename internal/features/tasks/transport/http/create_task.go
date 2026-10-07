package tasks_transport_http

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type CreateTaskRequest struct {
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	AuthorUserID uuid.UUID `json:"author_user_id" validate:"required"`
}

func (req *CreateTaskRequest) toDomain() domain.TaskCreate {
	return domain.TaskCreate{
		Title:        req.Title,
		Description:  req.Description,
		AuthorUserID: req.AuthorUserID,
	}
}

func (h *TasksHTTPHandler) CreateTask(c *core_http.Context) error {
	var req CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	task, err := h.tasksService.CreateTask(c.Context(), req.toDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, taskToResponse(task))
}
