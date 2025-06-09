package grpc

import (
	"context"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
)

type CustomerService struct {
	client protogen.MainClient
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetUser(ctx context.Context, req *protogen.GetUserRequest) (*protogen.UserResponse, error) {
	return s.client.GetUser(ctx, req)
}

func (s *CustomerService) CreateUser(ctx context.Context, req *protogen.CreateUserRequest) (*protogen.Empty, error) {
	return s.client.CreateUser(ctx, req)
}