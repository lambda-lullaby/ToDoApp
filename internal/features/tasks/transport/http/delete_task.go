package tasks_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *TasksHTTPHandler) DeleteTask(c *core_http.Context) error {
	id, err := parsePathID(c)
	if err != nil {
		return err
	}

	if err := h.tasksService.DeleteTask(c.Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
