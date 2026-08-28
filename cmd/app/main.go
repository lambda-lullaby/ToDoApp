package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/lambda-lullaby/ToDoApp/internal/core/config"
	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
	core_postgres_pgx "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool/pgx"
	core_http_middleware "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/middleware"
	core_http_server "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/server"

	users_postgres_repository "github.com/lambda-lullaby/ToDoApp/internal/features/users/repository/postgres"
	users_service "github.com/lambda-lullaby/ToDoApp/internal/features/users/service"
	users_transport_http "github.com/lambda-lullaby/ToDoApp/internal/features/users/transport/http"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.NewConfigMust()
	logger := core_logger.New(zapcore.InfoLevel)
	defer logger.Sync()

	postgresPool, err := core_postgres_pgx.New(ctx, cfg.Postgres.DSN(), cfg.Postgres.OpTimeout)
	if err != nil {
		logger.Error("connect to postgres", zap.Error(err))
		panic(err)
	}
	defer postgresPool.Close()

	usersRepository := users_postgres_repository.NewUsersRepository(postgresPool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	httpServer := core_http_server.NewHTTPServer(
		cfg.HTTP.Port,
		logger,
		core_http_middleware.CORS,
		core_http_middleware.RequestID,
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace,
		core_http_middleware.Panic,
	)

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.AddRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	httpServer.RegisterRoutes(core_http_server.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	})

	if err := httpServer.Run(ctx, cfg.HTTP.ShutdownTimeout); err != nil {
		logger.Error("http server run", zap.Error(err))
		panic(err)
	}
}
