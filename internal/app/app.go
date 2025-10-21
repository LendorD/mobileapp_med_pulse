package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/http/handlers"
	httpClient "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/http/onec"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/auth"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/swagger"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/workers"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
			func(cfg *config.Config) string { return cfg.JWTSecret },
		),
		LoggingModule,
		OneCModule,
		WebsocketModule,
		RepositoryModule,
		ServiceModule,
		UsecaseModule,
		HttpServerModule,
	)
}

func ProvideLoggers(cfg *config.Config) *logging.Logger {
	loggerCfg := &logging.Config{
		Enabled:    cfg.Logging.Enable,
		Level:      cfg.Logging.Level,
		LogsDir:    cfg.Logging.LogsDir,
		SavingDays: IntToUint(cfg.Logging.SavingDays),
	}

	logger := logging.NewLogger(loggerCfg, "APP", cfg.App.Version)
	return logger
}

var LoggingModule = fx.Module("logging_module",
	fx.Provide(
		ProvideLoggers,
	),
	fx.Invoke(func(l *logging.Logger) {
		l.Info("Logging system initialized")
	}),
)

func InvokeHttpServer(lc fx.Lifecycle, h http.Handler) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: h,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Close()
		},
	})
}

func NewSwaggerConfig(cfg *config.Config) *swagger.Config {
	return &swagger.Config{
		Enabled: true,
		Path:    "/swagger",
	}
}

var HttpServerModule = fx.Module("http_server_module",
	fx.Provide(
		NewSwaggerConfig,
		handlers.NewHandler,
		handlers.NewWebsocketHandler,
		handlers.ProvideRouter,
	),
	fx.Invoke(InvokeHttpServer),
)

var ServiceModule = fx.Module("service_module",
	fx.Provide(services.NewService),
)

var RepositoryModule = fx.Module("postgres_module",
	fx.Provide(repositories.NewRepository),
)

var UsecaseModule = fx.Module("usecases_module",
	fx.Provide(
		usecases.NewUsecases,
		usecases.NewAuthUsecase,
	),
)

func ProvideStdLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

var AuthModule = fx.Module("auth_module",
	fx.Provide(
		auth.NewAuthRepository,
		func(cfg *config.Config) string {
			if cfg == nil {
				panic("config is nil")
			}
			return cfg.JWTSecret
		},
		usecases.NewAuthUsecase,
	),
)

func ProvideOneCClient(cfg *config.Config) interfaces.OneCClient {
	return httpClient.NewOneCClient(cfg.OneC)
}

var OneCModule = fx.Module("onec_module",
	fx.Provide(
		ProvideOneCClient,
		usecases.NewOneCWebhookUsecase,
		ProvidePatientSyncWorker,
	),
)

var WebsocketModule = fx.Module("websocket_module",
	fx.Provide(ProvideStdLogger,
		websocket.NewHub,
	),
	fx.Invoke(websocket.InvokeHub),
)

func ProvidePatientSyncWorker(lc fx.Lifecycle, uc *usecases.OneCPatientUsecase, cfg *config.Config) *workers.PatientSyncWorker {
	interval := time.Minute * 5
	// if cfg.PatientSyncInterval > 0 {
	// 	interval = time.Duration(cfg.PatientSyncInterval) * time.Second
	// }

	worker := &workers.PatientSyncWorker{
		Usecase:  uc,
		Interval: interval,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go worker.Start(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			worker.Stop()
			return nil
		},
	})

	return worker
}

// TODO: Может быть вынести в services
func IntToUint(c int) uint {
	if c < 0 {
		panic([2]any{"a negative number", c})
	}
	return uint(c)
}
