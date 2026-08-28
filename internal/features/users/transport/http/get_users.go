package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limit, err := parseOptionalIntQuery(r, "limit")
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}
	offset, err := parseOptionalIntQuery(r, "offset")
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	users, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	responses := make([]UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, userToResponse(u))
	}
	core_http.RespondJSON(w, http.StatusOK, responses)
}
