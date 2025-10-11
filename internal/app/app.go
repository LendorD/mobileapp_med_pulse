package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/handlers"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories"
	"github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/auth"
	onecRepo "github.com/AlexanderMorozov1919/mobileapp/internal/adapters/repositories/onec"
	"github.com/AlexanderMorozov1919/mobileapp/internal/cache"
	"github.com/AlexanderMorozov1919/mobileapp/internal/config"
	"github.com/AlexanderMorozov1919/mobileapp/internal/external/onec"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/logging"
	"github.com/AlexanderMorozov1919/mobileapp/internal/middleware/swagger"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services"
	"github.com/AlexanderMorozov1919/mobileapp/internal/services/websocket"
	"github.com/AlexanderMorozov1919/mobileapp/internal/usecases"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func New() *fx.App {
	return fx.New(
		fx.Provide(
			config.LoadConfig,
			func(cfg *config.Config) string { return cfg.JWTSecret },
		),
		LoggingModule,
		CacheModule,
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

// В app/app.go

func ProvideRedisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}

func ProvideRedisCache(client *redis.Client) *cache.RedisCache {
	// Проверка подключения
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	return &cache.RedisCache{Client: client} // ← обнови RedisCache, чтобы он хранил *redis.Client
}

var CacheModule = fx.Module("cache_module",
	fx.Provide(ProvideRedisClient),
	fx.Provide(ProvideRedisCache),
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

// Swagger-конфигуратор
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

func ProvideOneCClient(cfg *config.Config) *onec.Client {
	return onec.NewClient(cfg.OneC)
}

var OneCModule = fx.Module("onec_module",
	fx.Provide(ProvideOneCClient),
	fx.Provide(onecRepo.NewRedisOneCCacheRepository),
	fx.Provide(usecases.NewOneCWebhookUsecase),
)

var WebsocketModule = fx.Module("websocket_module",
	fx.Provide(ProvideStdLogger),
	fx.Provide(websocket.NewHub),
	fx.Invoke(websocket.InvokeHub),
)

// TODO: Может быть вынести в services
func IntToUint(c int) uint {
	if c < 0 {
		panic([2]any{"a negative number", c})
	}
	return uint(c)
}
