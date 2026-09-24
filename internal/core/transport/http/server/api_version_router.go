package core_http_server

import (
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
)

type APIVersion string

const ApiVersion1 APIVersion = "v1"

type APIVersionRouter struct {
	version    APIVersion
	middleware []core_http.MiddlewareFunc
	routes     []Route
}

func NewAPIVersionRouter(version APIVersion, middleware ...core_http.MiddlewareFunc) *APIVersionRouter {
	return &APIVersionRouter{version: version, middleware: middleware}
}

func (r *APIVersionRouter) AddRoutes(routes ...Route) {
	r.routes = append(r.routes, routes...)
}
