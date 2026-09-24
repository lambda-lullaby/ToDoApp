package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *UsersHTTPHandler) GetUsers(c *core_http.Context) error {
	limit, err := parseOptionalIntQuery(c, "limit")
	if err != nil {
		return err
	}
	offset, err := parseOptionalIntQuery(c, "offset")
	if err != nil {
		return err
	}

	users, err := h.usersService.GetUsers(c.Context(), limit, offset)
	if err != nil {
		return err
	}

	responses := make([]UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, userToResponse(u))
	}
	return c.JSON(http.StatusOK, responses)
}
