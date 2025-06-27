package services

import (
	"context"
	"log/slog"

	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/repositories"
)

type UserService interface {
	GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error)
	StoreUserSession(ctx context.Context, user *entities.User) error
}

type userService struct {
	userRedisRepository repositories.UserRedisRepository
	logger              *slog.Logger
}

func NewUserService(userRedisRepository repositories.UserRedisRepository) UserService {
	return &userService{
		userRedisRepository: userRedisRepository,
		logger:              slog.Default().With("component", "user_service"),
	}
}

func (s *userService) GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error) {
	s.logger.Info("Getting user by session ID from repository",
		"session_id", sessionID,
		"method", "GetUserBySessionID")

	user, err := s.userRedisRepository.GetUserBySessionID(ctx, sessionID)
	if err != nil {
		s.logger.Error("Failed to get user by session ID from repository",
			"session_id", sessionID,
			"error", err.Error(),
			"method", "GetUserBySessionID")
		return nil, err
	}

	if user != nil {
		s.logger.Info("Successfully retrieved user by session ID",
			"session_id", sessionID,
			"user_id", user.ID,
			"username", user.Username,
			"method", "GetUserBySessionID")
	} else {
		s.logger.Info("No user found for session ID",
			"session_id", sessionID,
			"method", "GetUserBySessionID")
	}

	return user, nil
}

func (s *userService) StoreUserSession(ctx context.Context, user *entities.User) error {
	s.logger.Info("Storing user session in repository",
		"user_id", user.ID,
		"username", user.Username,
		"method", "StoreUserSession")

	err := s.userRedisRepository.StoreUserSession(ctx, user)
	if err != nil {
		s.logger.Error("Failed to store user session in repository",
			"user_id", user.ID,
			"username", user.Username,
			"error", err.Error(),
			"method", "StoreUserSession")
		return err
	}

	s.logger.Info("Successfully stored user session",
		"user_id", user.ID,
		"username", user.Username,
		"method", "StoreUserSession")

	return nil
}
