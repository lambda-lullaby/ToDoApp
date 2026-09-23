package users_transport_http

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func (h *UsersHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := parsePathID(r)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	err = h.usersService.DeleteUser(ctx, id)
	core_http.RespondEmpty(ctx, w, err)
}
