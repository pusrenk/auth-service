package handlers

import (
	"context"
	"log/slog"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/services"
)

type UserHandler struct {
	protogen.UnimplementedMainServer
	userService services.UserService
	logger      *slog.Logger
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		logger:      slog.Default().With("component", "user_handler"),
	}
}

func (h *UserHandler) GetUserBySessionID(ctx context.Context, req *protogen.GetUserBySessionIDRequest) (*protogen.UserResponse, error) {
	h.logger.Info("Getting user by session ID",
		"session_id", req.Id,
		"method", "GetUserBySessionID")

	user, err := h.userService.GetUserBySessionID(ctx, req.Id)
	if err != nil {
		h.logger.Error("Failed to get user by session ID",
			"session_id", req.Id,
			"error", err.Error(),
			"method", "GetUserBySessionID")
		return nil, err
	}

	h.logger.Info("Successfully retrieved user by session ID",
		"session_id", req.Id,
		"user_id", user.ID,
		"username", user.Username,
		"method", "GetUserBySessionID")

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
	h.logger.Info("Storing user session",
		"user_id", req.User.Id,
		"username", req.User.Username,
		"email", req.User.Email,
		"method", "StoreUserSession")

	user := &entities.User{
		ID:       req.User.Id,
		Username: req.User.Username,
		Email:    req.User.Email,
		Password: req.User.Password, // TODO: Hash the password
		Role:     req.User.Role,
	}

	err := h.userService.StoreUserSession(ctx, user)
	if err != nil {
		h.logger.Error("Failed to store user session",
			"user_id", req.User.Id,
			"username", req.User.Username,
			"email", req.User.Email,
			"error", err.Error(),
			"method", "StoreUserSession")
		return nil, err
	}

	h.logger.Info("Successfully stored user session",
		"user_id", req.User.Id,
		"username", req.User.Username,
		"method", "StoreUserSession")

	return &protogen.Empty{}, nil
}
