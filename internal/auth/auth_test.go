//go:generate mockgen -source=./auth.go -destination=./test_mock.go -package=auth_test
package auth_test

import (
	"avitotestgo2024/internal/auth"
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestHasPermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockRedisClient(ctrl)
	authInstance := auth.Auth{Base: mockClient, Logger: slog.Default()}

	t.Run("TokenExistsWithPermission", func(t *testing.T) {
		mockClient.EXPECT().Get(context.Background(), "valid_token").Return(redis.NewStringResult("1", nil))

		hasPermission, err := authInstance.HasPermission("valid_token", 1)

		assert.NoError(t, err)
		assert.True(t, hasPermission)
	})

	t.Run("TokenExistsWithoutPermission", func(t *testing.T) {
		mockClient.EXPECT().Get(context.Background(), "valid_token").Return(redis.NewStringResult("2", nil))

		hasPermission, err := authInstance.HasPermission("valid_token", 1)

		assert.EqualError(t, err, "403")
		assert.False(t, hasPermission)
	})

	t.Run("TokenDoesNotExist", func(t *testing.T) {
		mockClient.EXPECT().Get(context.Background(), "invalid_token").Return(redis.NewStringResult("", redis.Nil))

		_, err := authInstance.HasPermission("invalid_token", 1)

		assert.EqualError(t, err, "401")
	})

	t.Run("RedisError", func(t *testing.T) {
		mockClient.EXPECT().Get(context.Background(), "error_token").Return(redis.NewStringResult("", errors.New("mock error")))

		_, err := authInstance.HasPermission("error_token", 1)

		assert.Error(t, err)
	})
}

func TestCreateUserTokentWithPermission(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockRedisClient(ctrl)
	authInstance := auth.Auth{Base: mockClient, Logger: slog.Default()}

	t.Run("Success", func(t *testing.T) {
		mockClient.EXPECT().Set(context.Background(), "valid_token", 1, gomock.Any()).Return(redis.NewStatusResult("", nil))

		err := authInstance.CreateUserTokentWithPermission("valid_token", 1)

		assert.NoError(t, err)
	})

	t.Run("RedisError", func(t *testing.T) {
		mockClient.EXPECT().Set(context.Background(), "error_token", 1, gomock.Any()).Return(redis.NewStatusResult("", errors.New("mock error")))

		err := authInstance.CreateUserTokentWithPermission("error_token", 1)

		assert.Error(t, err)
	})
}
