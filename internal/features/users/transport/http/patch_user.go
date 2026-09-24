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

func (h *UsersHTTPHandler) PatchUser(c *core_http.Context) error {
	id, err := parsePathID(c)
	if err != nil {
		return err
	}

	var req PatchUserRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	user, err := h.usersService.PatchUser(c.Context(), id, req.toDomain())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userToResponse(user))
}
