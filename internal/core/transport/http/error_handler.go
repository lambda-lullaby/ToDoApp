package core_http

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	core_errors "github.com/lambda-lullaby/ToDoApp/internal/core/errors"
	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
)

type errorResponse struct {
	Error string `json:"error"`
}

func DefaultErrorHandler(err error, c *Context) {
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
		core_logger.FromContext(c.Context()).Error("unhandled error", zap.Error(err))
	}

	_ = c.JSON(statusCode, errorResponse{Error: message})
}
