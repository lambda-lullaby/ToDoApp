package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := parsePathID(r)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	user, err := h.usersService.GetUser(ctx, id)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	core_http.RespondJSON(w, http.StatusOK, userToResponse(user))
}
