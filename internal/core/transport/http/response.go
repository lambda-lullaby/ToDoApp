package core_http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
)

type errorResponse struct {
	Error string `json:"error"`
}

func RespondJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body == nil {
		return
	}
	_ = json.NewEncoder(w).Encode(body)
}

func RespondNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func RespondError(ctx context.Context, w http.ResponseWriter, err error) {
	var statusCode int
	message := err.Error()

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
	case errors.Is(err, core_errors.ErrNotFound):
		statusCode = http.StatusNotFound
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
		message = "internal server error"
		core_logger.FromContext(ctx).Error("unhandled error", zap.Error(err))
	}

	RespondJSON(w, statusCode, errorResponse{Error: message})
}
