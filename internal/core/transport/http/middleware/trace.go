package core_http_middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
)

func Trace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := NewResponseWriter(w)

		next.ServeHTTP(rw, r)

		core_logger.FromContext(r.Context()).Info("request handled",
			zap.String("method", r.Method),
			zap.Int("status_code", rw.StatusCode),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
