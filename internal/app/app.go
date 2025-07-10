package app

import (
	"context"
	"net/http"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/handlers"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/auth"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
			// Добавляем провайдер для JWT secret
			func(cfg *config.Config) string { return cfg.JWTSecret },
		),
		LoggingModule,
		RepositoryModule,
		ServiceModule,
		UsecaseModule,
		HttpServerModule,
	)
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

func InvokeHttpServer(lc fx.Lifecycle, h http.Handler) {
	server := &http.Server{
		Addr:    ":8080", // Упрощаем - используем хардкод порта
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
	fx.Provide(
		usecases.NewUsecases,
		usecases.NewAuthUsecase,
	),
)

/* -------------------------------------------- */

func IntToUint(c int) uint {
	if c < 0 {
		panic([2]any{"a negative number", c})
	}

	return uint(c)
}
