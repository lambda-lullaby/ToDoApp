package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}
