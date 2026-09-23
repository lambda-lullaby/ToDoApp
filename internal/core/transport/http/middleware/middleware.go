package core_http_middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) == 0 {
		return h
	}
	return ChainMiddleware(m[len(m)-1](h), m[:len(m)-1]...)
}
