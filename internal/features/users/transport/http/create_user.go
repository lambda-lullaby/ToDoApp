package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

func (h *UsersHTTPHandler) CreateUser(c *core_http.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.usersService.CreateUser(c.Context(), req.FullName, req.PhoneNumber)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userToResponse(user))
}
