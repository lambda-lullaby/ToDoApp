package core_http_middleware

import (
	"fmt"

	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func Panic(next core_http.HandlerFunc) core_http.HandlerFunc {
	return func(c *core_http.Context) (err error) {
		defer func() {
			if rec := recover(); rec != nil {
				core_logger.FromContext(c.Context()).Error("panic recovered", zap.Any("panic", rec), zap.Stack("stack"))
				err = fmt.Errorf("panic recovered: %v", rec)
			}
		}()
		return next(c)
	}
}
