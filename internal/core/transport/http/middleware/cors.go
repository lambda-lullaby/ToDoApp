package core_http_middleware

import (
	"net/http"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

func CORS(next core_http.HandlerFunc) core_http.HandlerFunc {
	return func(c *core_http.Context) error {
		header := c.Response().Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")

		if c.Request().Method == http.MethodOptions {
			return c.NoContent(http.StatusNoContent)
		}

		return next(c)
	}
}
