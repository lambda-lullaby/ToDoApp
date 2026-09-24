package core_http_middleware

import (
	"context"

	"github.com/google/uuid"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type requestIDCtxKey struct{}

const RequestIDHeader = "X-Request-ID"

func RequestID(next core_http.HandlerFunc) core_http.HandlerFunc {
	return func(c *core_http.Context) error {
		requestID := c.Request().Header.Get(RequestIDHeader)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Response().Header().Set(RequestIDHeader, requestID)
		c.SetContext(context.WithValue(c.Context(), requestIDCtxKey{}, requestID))
		return next(c)
	}
}

func RequestIDFromContext(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDCtxKey{}).(string)
	return requestID
}
