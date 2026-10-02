package users_transport_http

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type UserResponse struct {
	ID          string  `json:"id"`
	Version     int64   `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userToResponse(u domain.User) UserResponse {
	return UserResponse{
		ID:          u.ID.String(),
		Version:     u.Version,
		FullName:    u.FullName,
		PhoneNumber: u.PhoneNumber,
	}
}

func parsePathID(c *core_http.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse `id` path parameter: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return id, nil
}

func parseOptionalIntQuery(c *core_http.Context, key string) (*int, error) {
	raw := c.QueryParam(key)
	if raw == "" {
		return nil, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return nil, fmt.Errorf("parse `%s` query parameter: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return &value, nil
}
