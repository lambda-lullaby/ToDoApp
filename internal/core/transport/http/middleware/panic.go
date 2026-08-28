package core_http_middleware

import (
	"net/http"

	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
)

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				core_logger.FromContext(r.Context()).Error("panic recovered", zap.Any("panic", rec))
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
