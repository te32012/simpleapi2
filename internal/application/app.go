package application

import (
	"avitotestgo2024/internal/auth"
	"avitotestgo2024/internal/database"
	"avitotestgo2024/internal/middlware"
	"avitotestgo2024/internal/service"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() {
	log := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(log)
	slog.Info("starting application")
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRESQL_USER"), os.Getenv("POSTGRESQL_PASSWORD"), os.Getenv("POSTGRESQL_HOST"), os.Getenv("POSTGRESQL_PORT"), os.Getenv("POSTGRESQL_BASE")))
	if err != nil {
		slog.Error("error connecting to database", slog.Any("error", err.Error()))
		return
	}
	defer pool.Close()
	err = pool.Ping(context.Background())
	if err != nil {
		slog.Error("error ping to database", slog.Any("error", err.Error()))
		return
	}
	base := database.NewDatabase(pool, slog.Default())
	slog.Info("sucsesfull connect to database")
	auth := auth.NewAuth("redis:6379", os.Getenv("REDIS_USERNAME"), os.Getenv("REDIS_PASSWORD"), slog.Default())
	err = auth.CreateUserTokentWithPermission("admin", 2)
	if err != nil {
		slog.Error("error save admin token in redis", slog.Any("error", err.Error()))
		return
	}
	err = auth.CreateUserTokentWithPermission("user", 1)
	if err != nil {
		slog.Error("error save user token in redis", slog.Any("error", err.Error()))
		return
	}
	slog.Info("sucsesfull connect to redis")
	service := service.NewService(base, slog.Default(), auth)
	middlvare := middlware.NewServer("application", "2024", auth, service, slog.Default())
	go func() {
		slog.Info("starting router")
		middlvare.Run()
		slog.Info("stop router")
	}()
	slog.Info("success starting application")

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	slog.Info("application has been shut down")
}
