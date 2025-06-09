package services

import (
	"context"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/handlers/grpc/customerService"
	"github.com/pusrenk/auth-service/internal/user/repositories/redis"
	"github.com/pusrenk/auth-service/pkg/helpers"
)

type UserService interface {
	GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error)
	StoreUser(ctx context.Context, user *entities.User) error
}

type userService struct {
	customerService customerService.CustomerService
	userRedisRepository redis.UserRedisRepository 
}

func NewUserService(userRedisRepository redis.UserRedisRepository) UserService {
	return &userService{userRedisRepository: userRedisRepository}
}

func (s *userService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	return s.customerService.GetUser(ctx, &protogen.GetUserRequest{Id: id})
}

func (s *userService) CreateUser(ctx context.Context, user *entities.User) error {
	return s.customerService.CreateUser(ctx, &protogen.CreateUserRequest{User: &protogen.BaseUser{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Role:     helpers.UserRole,
	}})
}

func (s *userService) GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error) {
	return s.userRedisRepository.GetUserBySessionID(ctx, sessionID)
}

func (s *userService) StoreUserSession(ctx context.Context, user *entities.User) error {
	return s.userRedisRepository.StoreUserSession(ctx, user)
}