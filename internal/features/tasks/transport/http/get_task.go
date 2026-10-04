package tasks_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *TasksHTTPHandler) GetTask(c *core_http.Context) error {
	id, err := parsePathID(c)
	if err != nil {
		return err
	}

	task, err := h.tasksService.GetTask(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, taskToResponse(task))
}
