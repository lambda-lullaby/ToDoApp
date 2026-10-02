package core_http_middleware

import (
	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func Logger(base *zap.Logger) core_http.MiddlewareFunc {
	return func(next core_http.HandlerFunc) core_http.HandlerFunc {
		return func(c *core_http.Context) error {
			logger := base.With(
				zap.String("request_id", RequestIDFromContext(c.Context())),
				zap.String("url", c.Request().URL.String()),
			)
			c.SetContext(core_logger.ToContext(c.Context(), logger))
			return next(c)
		}
	}
}
