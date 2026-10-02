package core_http

type HandlerFunc func(c *Context) error

type MiddlewareFunc func(next HandlerFunc) HandlerFunc
