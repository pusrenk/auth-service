package services

import (
	"context"

	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/repositories"
)

type UserService interface {
	GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error)
	StoreUserSession(ctx context.Context, user *entities.User) error
}

type userService struct {
	userRedisRepository repositories.UserRedisRepository
}

func NewUserService(userRedisRepository repositories.UserRedisRepository) UserService {
	return &userService{userRedisRepository: userRedisRepository}
}

func (s *userService) GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error) {
	return s.userRedisRepository.GetUserBySessionID(ctx, sessionID)
}

func (s *userService) StoreUserSession(ctx context.Context, user *entities.User) error {
	return s.userRedisRepository.StoreUserSession(ctx, user)
}
