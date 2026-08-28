package core_http_middleware

import (
	"net/http"

	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
)

func Logger(base *zap.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger := base.With(
				zap.String("request_id", RequestIDFromContext(r.Context())),
				zap.String("url", r.URL.String()),
			)
			ctx := core_logger.ToContext(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
