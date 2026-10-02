package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *UsersHTTPHandler) GetUser(c *core_http.Context) error {
	id, err := parsePathID(c)
	if err != nil {
		return err
	}

	user, err := h.usersService.GetUser(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userToResponse(user))
}
