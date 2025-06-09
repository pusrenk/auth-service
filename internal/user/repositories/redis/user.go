package redis

import (
	"context"
	"encoding/json"

	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/pkg/helpers"
	"github.com/redis/go-redis/v9"
)

type UserRedisRepository interface {
	GetUserBySessionID(ctx context.Context, id string) (*entities.User, error)
	StoreUserSession(ctx context.Context, user *entities.User) error
}

type userRedisRepository struct {
	rdb *redis.Client
}

func NewUserRedisRepository(rdb *redis.Client) UserRedisRepository {
	return &userRedisRepository{rdb: rdb}
}

func (r *userRedisRepository) GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error) {
	data, err := r.rdb.Get(ctx, sessionID).Result()
	if err != nil {
		return nil, err
	}

	var user entities.User

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRedisRepository) StoreUserSession(ctx context.Context, user *entities.User) error {
	sessionID := helpers.GenerateSessionID()
	jsonData, err := json.Marshal(user)

	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, sessionID, jsonData, helpers.SessionExpiry).Err()
}
