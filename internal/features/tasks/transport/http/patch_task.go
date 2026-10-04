package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
	core_http_types "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (req *PatchTaskRequest) Validate() error {
	if req.Title.Set && req.Title.Value == nil {
		return fmt.Errorf("`title` can't be null")
	}
	if req.Completed.Set && req.Completed.Value == nil {
		return fmt.Errorf("`completed` can't be null")
	}
	return nil
}

func (req *PatchTaskRequest) toDomain() domain.TaskPatch {
	return domain.TaskPatch{
		Title:       req.Title.ToDomain(),
		Description: req.Description.ToDomain(),
		Completed:   req.Completed.ToDomain(),
	}
}

func (h *TasksHTTPHandler) PatchTask(c *core_http.Context) error {
	id, err := parsePathID(c)
	if err != nil {
		return err
	}

	var req PatchTaskRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	task, err := h.tasksService.PatchTask(c.Context(), id, req.toDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, taskToResponse(task))
}
