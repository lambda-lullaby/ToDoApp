package core_http_middleware

import (
	"time"

	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func Trace(next core_http.HandlerFunc) core_http.HandlerFunc {
	return func(c *core_http.Context) error {
		start := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		core_logger.FromContext(c.Context()).Info("request handled",
			zap.String("method", c.Request().Method),
			zap.Int("status_code", c.Response().Status),
			zap.Duration("duration", time.Since(start)),
		)
		return err
	}
}
