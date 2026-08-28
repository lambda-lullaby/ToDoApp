package users_transport_http

import (
	"fmt"
	"net/http"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
	core_http_types "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (req *PatchUserRequest) Validate() error {
	if req.FullName.Set && req.FullName.Value == nil {
		return fmt.Errorf("`full_name` can't be null")
	}
	return nil
}

func (req *PatchUserRequest) toDomain() domain.UserPatch {
	return domain.UserPatch{
		FullName:    req.FullName.ToDomain(),
		PhoneNumber: req.PhoneNumber.ToDomain(),
	}
}

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := parsePathID(r)
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	var req PatchUserRequest
	if err := core_http.DecodeAndValidateRequest(r, &req); err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	user, err := h.usersService.PatchUser(ctx, id, req.toDomain())
	if err != nil {
		core_http.RespondError(ctx, w, err)
		return
	}

	core_http.RespondJSON(w, http.StatusOK, userToResponse(user))
}
