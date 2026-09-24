package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
	core_http_middleware "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/middleware"
)

type HTTPServer struct {
	mux        *http.ServeMux
	server     *http.Server
	middleware []core_http.MiddlewareFunc
	logger     *zap.Logger
}

func NewHTTPServer(port int, logger *zap.Logger, globalMiddleware ...core_http.MiddlewareFunc) *HTTPServer {
	mux := http.NewServeMux()
	return &HTTPServer{
		mux:        mux,
		server:     &http.Server{Addr: ":" + strconv.Itoa(port), Handler: mux},
		middleware: globalMiddleware,
		logger:     logger,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		for _, route := range router.routes {
			handler := core_http_middleware.ChainMiddleware(route.Handler, route.Middleware...)
			handler = core_http_middleware.ChainMiddleware(handler, router.middleware...)

			pattern := fmt.Sprintf("%s /api/%s%s", route.Method, router.version, route.Path)
			s.mount(pattern, handler)
		}
	}
}

func (s *HTTPServer) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		handler := core_http_middleware.ChainMiddleware(route.Handler, route.Middleware...)

		pattern := route.Path
		if route.Method != "" {
			pattern = route.Method + " " + route.Path
		}
		s.mount(pattern, handler)
	}
}

func (s *HTTPServer) RegisterSwagger(pathPrefix string, handler http.Handler) {
	s.mount(pathPrefix, func(c *core_http.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}

func (s *HTTPServer) mount(pattern string, handler core_http.HandlerFunc) {
	handler = core_http_middleware.ChainMiddleware(handler, s.middleware...)

	s.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		c := core_http.NewContext(w, r)
		if err := handler(c); err != nil {
			c.Error(err)
		}
	})
}

func (s *HTTPServer) Run(ctx context.Context, shutdownTimeout time.Duration) error {
	errCh := make(chan error, 1)
	go func() {
		s.logger.Info("http server listening", zap.String("addr", s.server.Addr))
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("listen and serve: %w", err)
			return
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	s.logger.Info("http server shutting down")
	if err := s.server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return <-errCh
}
