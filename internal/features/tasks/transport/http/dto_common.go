package tasks_transport_http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type TaskResponse struct {
	ID           string     `json:"id"`
	Version      int64      `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID string     `json:"author_user_id"`
}

func taskToResponse(t domain.Task) TaskResponse {
	return TaskResponse{
		ID:           t.ID.String(),
		Version:      t.Version,
		Title:        t.Title,
		Description:  t.Description,
		Completed:    t.Completed,
		CreatedAt:    t.CreatedAt,
		CompletedAt:  t.CompletedAt,
		AuthorUserID: t.AuthorUserID.String(),
	}
}

func parsePathID(c *core_http.Context) (uuid.UUID, error) {
	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse `id` path parameter: %v: %w", err, core_errors.ErrInvalidArgument)
	}
	return id, nil
}

func parseIntQuery(c *core_http.Context, key string, defaultValue int) (int, error) {
	raw := c.QueryParam(key)
	if raw == "" {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parse `%s` query parameter: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return value, nil
}

func parseOptionalUUIDQuery(c *core_http.Context, key string) (*uuid.UUID, error) {
	raw := c.QueryParam(key)
	if raw == "" {
		return nil, nil
	}
	value, err := uuid.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("parse `%s` query parameter: %v: %w", key, err, core_errors.ErrInvalidArgument)
	}
	return &value, nil
}
