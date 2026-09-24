package core_http

import "net/http"

type Response struct {
	http.ResponseWriter
	Status    int
	Committed bool
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{ResponseWriter: w, Status: http.StatusOK}
}

func (r *Response) WriteHeader(status int) {
	if r.Committed {
		return
	}
	r.Status = status
	r.Committed = true
	r.ResponseWriter.WriteHeader(status)
}

func (r *Response) Write(b []byte) (int, error) {
	if !r.Committed {
		r.WriteHeader(http.StatusOK)
	}
	return r.ResponseWriter.Write(b)
}
