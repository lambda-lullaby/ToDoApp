package tasks_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

const (
	defaultTasksLimit  = 100
	defaultTasksOffset = 0
)

func (h *TasksHTTPHandler) GetTasks(c *core_http.Context) error {
	userID, err := parseOptionalUUIDQuery(c, "user_id")
	if err != nil {
		return err
	}
	limit, err := parseIntQuery(c, "limit", defaultTasksLimit)
	if err != nil {
		return err
	}
	offset, err := parseIntQuery(c, "offset", defaultTasksOffset)
	if err != nil {
		return err
	}

	tasks, err := h.tasksService.GetTasks(c.Context(), userID, limit, offset)
	if err != nil {
		return err
	}

	responses := make([]TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		responses = append(responses, taskToResponse(t))
	}
	return c.JSON(http.StatusOK, responses)
}
