package services

import (
	"context"

	"github.com/pusrenk/auth-service/internal/protobuf/customer-service/protogen"
	"github.com/pusrenk/auth-service/internal/user/entities"
	rpc "github.com/pusrenk/auth-service/internal/user/handlers/grpc/customer-service"
	"github.com/pusrenk/auth-service/internal/user/repositories/redis"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error)
	StoreUserSession(ctx context.Context, user *entities.User) error
}

type userService struct {
	customerService     *rpc.CustomerService
	userRedisRepository redis.UserRedisRepository
}

func NewUserService(
	customerSvc *rpc.CustomerService,
	userRedisRepository redis.UserRedisRepository,
) UserService {
	return &userService{
		customerService:     customerSvc,
		userRedisRepository: userRedisRepository,
	}
}

func (s *userService) GetUser(ctx context.Context, id string) (*entities.User, error) {
	res, err := s.customerService.GetUser(ctx, &protogen.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       res.User.Id,
		Username: res.User.Username,
		Email:    res.User.Email,
		Role:     res.User.Role,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, user *entities.User) error {
	_, err := s.customerService.CreateUser(ctx, &protogen.CreateUserRequest{
		User: &protogen.BaseUser{
			Id:       user.ID,
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		},
	})

	return err
}

func (s *userService) GetUserBySessionID(ctx context.Context, sessionID string) (*entities.User, error) {
	return s.userRedisRepository.GetUserBySessionID(ctx, sessionID)
}

func (s *userService) StoreUserSession(ctx context.Context, user *entities.User) error {
	return s.userRedisRepository.StoreUserSession(ctx, user)
}
