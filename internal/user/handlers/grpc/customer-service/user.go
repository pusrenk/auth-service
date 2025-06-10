package rpc

import (
	"context"

	"github.com/pusrenk/auth-service/internal/protobuf/customer-service/protogen"
)

type CustomerService struct {
	customerRpcClient protogen.CustomerRpcClient
}

func NewCustomerService(customerRpcClient protogen.CustomerRpcClient) *CustomerService {
	return &CustomerService{customerRpcClient: customerRpcClient}
}

func (s *CustomerService) GetUser(ctx context.Context, req *protogen.GetUserRequest) (*protogen.UserResponse, error) {
	return s.customerRpcClient.GetUser(ctx, req)
}

func (s *CustomerService) CreateUser(ctx context.Context, req *protogen.CreateUserRequest) (*protogen.Empty, error) {
	return s.customerRpcClient.CreateUser(ctx, req)
}
