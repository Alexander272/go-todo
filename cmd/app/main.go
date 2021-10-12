package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alexander272/go-todo/internal/config"
	delivery "github.com/Alexander272/go-todo/internal/delivery/http"
	"github.com/Alexander272/go-todo/internal/repository"
	"github.com/Alexander272/go-todo/internal/server"
	"github.com/Alexander272/go-todo/internal/service"
	"github.com/Alexander272/go-todo/pkg/auth"
	"github.com/Alexander272/go-todo/pkg/database/mongodb"
	"github.com/Alexander272/go-todo/pkg/database/redis"
	"github.com/Alexander272/go-todo/pkg/hash"
	"github.com/Alexander272/go-todo/pkg/logger"
	"github.com/joho/godotenv"
)

// @title Todo App
// @version 0.1
// @description API Server for Todo App

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	conf, err := config.Init("configs")
	if err != nil {
		logger.Fatalf("error initializing configs: %s", err.Error())
	}
	logger.Init(os.Stdout, conf.Environment)

	// Dependencies
	mongoClient, err := mongodb.NewClient(conf.Mongo.URI, conf.Mongo.User, conf.Mongo.Password)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}
	db := mongoClient.Database(conf.Mongo.Name)

	client, err := redis.NewRedisClient(redis.Config{
		Host:     conf.Redis.Host,
		Port:     conf.Redis.Port,
		DB:       conf.Redis.DB,
		Password: conf.Redis.Password,
	})
	if err != nil {
		logger.Fatalf("failed to initialize redis %s", err.Error())
	}

	hasher := hash.NewBcryptHasher(conf.Auth.Bcrypt.MinCost, conf.Auth.Bcrypt.DefaultCost, conf.Auth.Bcrypt.MaxCost)
	tokenManager, err := auth.NewManager(conf.Auth.JWT.Key)
	if err != nil {
		logger.Fatalf("failed to initialize token manager: %s", err.Error())
	}

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(db, client)
	services := service.NewServices(service.Deps{
		Repos:                  repos,
		Hasher:                 hasher,
		TokenManager:           tokenManager,
		AccessTokenTTL:         conf.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL:        conf.Auth.JWT.RefreshTokenTTL,
		Domain:                 conf.Http.Host,
		VerificationCodeLength: conf.Auth.VerificationCodeLength,
	})
	handlers := delivery.NewHandler(services)

	// HTTP Server
	srv := server.NewServer(conf, handlers.Init(conf))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Infof("Application started on port: %s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		logger.Errorf("error occured on db connection close: %s", err.Error())
	}
}
