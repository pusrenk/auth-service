package handlers

import (
	"context"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/services"
)

type UserHandler struct {
	protogen.UnimplementedMainServer
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUserBySessionID(ctx context.Context, req *protogen.GetUserBySessionIDRequest) (*protogen.UserResponse, error) {
	user, err := h.userService.GetUserBySessionID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &protogen.UserResponse{
		User: &protogen.BaseUser{
			Id:       user.ID,
			Username: user.Username,
			Password: user.Password,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, nil
}

func (h *UserHandler) StoreUserSession(ctx context.Context, req *protogen.StoreUserSessionRequest) (*protogen.Empty, error) {
	user := &entities.User{
		ID:       req.User.Id,
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password, // TODO: Hash the password
		Role:     req.User.Role,
	}

	err := h.userService.StoreUserSession(ctx, user)
	if err != nil {
		return nil, err
	}

	return &protogen.Empty{}, nil
}
