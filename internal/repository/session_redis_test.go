package repository

import (
	"context"
	"os"
	"testing"
	"time"

	redisDB "github.com/Alexander272/go-todo/pkg/database/redis"
	"github.com/Alexander272/go-todo/pkg/logger"
	"github.com/alicebob/miniredis/v2"
	"github.com/elliotchance/redismock/v8"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var client *redis.Client
var (
	key = "token"
	val = SessionData{
		UserId: "UserId",
		Email:  "Email",
		Role:   "Role",
		Ua:     "Ua",
		Ip:     "Ip",
		Exp:    time.Duration(5),
	}
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		logger.Fatalf("failed to open miniredis. error %s", err.Error())
	}
	defer mr.Close()

	client, err = redisDB.NewRedisClient(redisDB.Config{
		Host: mr.Host(),
		Port: mr.Port(),
	})
	if err != nil {
		logger.Fatalf("failed to initialize redis %s", err.Error())
	}

	code := m.Run()
	os.Exit(code)
}

func TestSessionRepository_Create(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Set", context.TODO(), key, val, val.Exp).Return(redis.NewStatusResult("", nil))

	r := NewSessionRepo(mock)
	err := r.CreateSession(context.TODO(), key, val)
	assert.NoError(t, err)
}
