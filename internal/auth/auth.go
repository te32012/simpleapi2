package auth

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type Auth struct {
	Base   RedisClient
	Logger *slog.Logger
}

func NewAuth(uri string, username, password string, logger *slog.Logger) *Auth {
	rdb := redis.NewClient(&redis.Options{
		Addr:     uri,
		Username: username,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	return &Auth{Base: rdb, Logger: logger}
}

func (a *Auth) HasPermission(s string, permission int) (bool, error) {
	a.Logger.Info("Проверка токента %s на наличие доступа уровня %d", s, permission)
	val, err := a.Base.Get(context.Background(), s).Int()
	if err != nil {
		a.Logger.Error("Токент %s не содержится в базе токентов")
		return false, errors.New("401")
	}
	if val == permission {
		a.Logger.Info("Токент %s имеет уровень доступа %d", s, permission)
		return true, nil
	}
	a.Logger.Error("Токент %s не имеет уровня доступа %d", s, permission)
	return false, errors.New("403")
}

func (a *Auth) CreateUserTokentWithPermission(s string, permission int) error {
	a.Logger.Info("Загрузка в базу токента %s с уровнем доступа %d", s, permission)
	err := a.Base.Set(context.Background(), s, permission, 0).Err()
	if err != nil {
		a.Logger.Error("Ошибка загрузки токента %s с уровнем доступа %d", s, permission)
		return err
	}
	a.Logger.Info("Токент %s с уровнем доступа %d загружен в базу", s, permission)
	return nil
}
