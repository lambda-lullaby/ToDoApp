package core_http_server

import (
	core_http_middleware "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/middleware"
)

type APIVersion string

const ApiVersion1 APIVersion = "v1"

type APIVersionRouter struct {
	version    APIVersion
	middleware []core_http_middleware.Middleware
	routes     []Route
}

func NewAPIVersionRouter(version APIVersion, middleware ...core_http_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{version: version, middleware: middleware}
}

func (r *APIVersionRouter) AddRoutes(routes ...Route) {
	r.routes = append(r.routes, routes...)
}
