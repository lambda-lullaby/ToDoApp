package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req CreateUserRequest
	if err := core_http.DecodeAndValidateRequest(r, &req); err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	user, err := h.usersService.CreateUser(ctx, req.FullName, req.PhoneNumber)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	core_http.RespondJSON(w, http.StatusCreated, userToResponse(user))
}
