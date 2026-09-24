package core_http

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	request  *http.Request
	response *Response
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{request: r, response: NewResponse(w)}
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Response() *Response {
	return c.response
}

func (c *Context) Context() context.Context {
	return c.request.Context()
}

func (c *Context) SetContext(ctx context.Context) {
	c.request = c.request.WithContext(ctx)
}

func (c *Context) PathParam(name string) string {
	return c.request.PathValue(name)
}

func (c *Context) QueryParam(name string) string {
	return c.request.URL.Query().Get(name)
}

func (c *Context) JSON(statusCode int, body any) error {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(statusCode)
	return json.NewEncoder(c.response).Encode(body)
}

func (c *Context) NoContent(statusCode int) error {
	c.response.WriteHeader(statusCode)
	return nil
}

func (c *Context) Error(err error) {
	if c.response.Committed {
		return
	}
	DefaultErrorHandler(err, c)
}
