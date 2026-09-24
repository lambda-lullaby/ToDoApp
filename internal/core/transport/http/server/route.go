package core_http_server

import (
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type Route struct {
	Method     string
	Path       string
	Handler    core_http.HandlerFunc
	Middleware []core_http.MiddlewareFunc
}
