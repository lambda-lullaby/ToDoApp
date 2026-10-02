package core_http_middleware

import (
	"context"
	"time"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func Timeout(d time.Duration) core_http.MiddlewareFunc {
	return func(next core_http.HandlerFunc) core_http.HandlerFunc {
		return func(c *core_http.Context) error {
			ctx, cancel := context.WithTimeout(c.Context(), d)
			defer cancel()

			c.SetContext(ctx)
			return next(c)
		}
	}
}
