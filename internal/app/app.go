package app

import (
	"context"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/handlers"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
		),

		// LoggingModule,

		RepositoryModule,
		ServiceModule,
		UsecaseModule,
		HttpServerModule,
	)
}

/*
func ProvideLoggers(cfg *config.Config) *logging.Logger {
	loggerCfg := logging.Config(
		cfg.Logging.Enable,
		cfg.Logging.Level,
		cfg.Logging.Format,
		cfg.Logging.LogsDir,
		IntToUint(cfg.Logging.SavingDays),
	)

	logger := logging.NewBaseLogger("VERSION", cfg.App.Version, loggerCfg, logging.WithDailyLogDelete())

	return logger
}

var LoggingModule = fx.Module("logging_module",
	fx.Provide(
		ProvideLoggers,
	),
	fx.Invoke(logging.InvokeBaseLogger),
)


*/
/* -------------------------------------------- */

func InvokeHttpServer(lc fx.Lifecycle, cfg *config.Config, h http.Handler) {
	server := &http.Server{
		Addr:    ":" + cfg.HTTPServer.Port,
		Handler: h,

		// Secutiry timeouts
		ReadHeaderTimeout: cfg.HTTPServer.ReadHeaderTimeout,
		ReadTimeout:       cfg.HTTPServer.ReadTimeout,
		WriteTimeout:      cfg.HTTPServer.WriteTimeout,
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := server.ListenAndServe(); err != nil {
						panic(err)
					}
				}()

				return nil
			},

			// Shutdown gracefully shuts down the server without interrupting any active connections.
			// Shutdown works by first closing all open listeners and then
			// waits indefinitely for all connections to return to idle before shutting down.
			OnStop: func(ctx context.Context) error {
				if err := server.Close(); err != nil {
					return err
				}

				return nil
			},
		},
	)
}

var HttpServerModule = fx.Module("http_server_module",
	fx.Provide(
		handlers.NewHandler,
		handlers.ProvideRouter,
	),

	fx.Invoke(InvokeHttpServer),
)

/* -------------------------------------------- */

var ServiceModule = fx.Module("service_module",
	fx.Provide(services.NewService),
)

/* -------------------------------------------- */

var RepositoryModule = fx.Module("postgres_module",
	fx.Provide(repositories.NewRepository),
)

/* -------------------------------------------- */

var UsecaseModule = fx.Module("usecases_module",
	fx.Provide(usecases.NewUsecases),
)

/* -------------------------------------------- */

func IntToUint(c int) uint {
	if c < 0 {
		panic([2]any{"a negative number", c})
	}

	return uint(c)
}
