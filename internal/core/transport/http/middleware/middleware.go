package core_http_middleware

import core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"

func ChainMiddleware(h core_http.HandlerFunc, m ...core_http.MiddlewareFunc) core_http.HandlerFunc {
	if len(m) == 0 {
		return h
	}
	return ChainMiddleware(m[len(m)-1](h), m[:len(m)-1]...)
}
