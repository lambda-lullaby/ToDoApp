package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/lambda-lullaby/ToDoApp/internal/core/config"
	"github.com/lambda-lullaby/ToDoApp/internal/core/domain"
	core_logger "github.com/lambda-lullaby/ToDoApp/internal/core/logger"
	core_postgres_pgx "github.com/lambda-lullaby/ToDoApp/internal/core/postgres/pool/pgx"
	core_http "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http"
	core_http_middleware "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/middleware"
	core_http_server "github.com/lambda-lullaby/ToDoApp/internal/core/transport/http/server"

	tasks_postgres_repository "github.com/lambda-lullaby/ToDoApp/internal/features/tasks/repository/postgres"
	tasks_service "github.com/lambda-lullaby/ToDoApp/internal/features/tasks/service"
	tasks_transport_http "github.com/lambda-lullaby/ToDoApp/internal/features/tasks/transport/http"
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

	time.Local = cfg.TimeZone
	logger.Info("time zone set", zap.String("time_zone", cfg.TimeZone.String()))

	domain.RegisterTaskValidation(domain.TaskLimits{
		TitleMinLength:       cfg.Task.TitleMinLength,
		TitleMaxLength:       cfg.Task.TitleMaxLength,
		DescriptionMinLength: cfg.Task.DescriptionMinLength,
		DescriptionMaxLength: cfg.Task.DescriptionMaxLength,
	})

	postgresPool, err := core_postgres_pgx.New(ctx, cfg.Postgres.DSN(), cfg.Postgres.OpTimeout)
	if err != nil {
		logger.Error("connect to postgres", zap.Error(err))
		panic(err)
	}
	defer postgresPool.Close()

	usersRepository := users_postgres_repository.NewUsersRepository(postgresPool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	tasksRepository := tasks_postgres_repository.NewTasksRepository(postgresPool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	httpServer := core_http_server.NewHTTPServer(
		cfg.HTTP.Port,
		logger,
		core_http_middleware.CORS,
		core_http_middleware.RequestID,
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace,
		core_http_middleware.Panic,
		core_http_middleware.Timeout(cfg.HTTP.RequestTimeout),
	)

	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.AddRoutes(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.AddRoutes(tasksTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	httpServer.RegisterRoutes(core_http_server.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(c *core_http.Context) error {
			return c.NoContent(http.StatusOK)
		},
	})

	if err := httpServer.Run(ctx, cfg.HTTP.ShutdownTimeout); err != nil {
		logger.Error("http server run", zap.Error(err))
		panic(err)
	}
}
