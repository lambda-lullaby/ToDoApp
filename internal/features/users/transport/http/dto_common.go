package users_transport_http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
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

func parsePathID(r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse `id` path parameter: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return id, nil
}

func parseOptionalIntQuery(r *http.Request, key string) (*int, error) {
	raw := r.URL.Query().Get(key)
	if raw == "" {
		return nil, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return nil, fmt.Errorf("parse `%s` query parameter: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return &value, nil
}
