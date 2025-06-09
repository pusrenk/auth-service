package grpc

import (
	"log"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/pkg/helpers"
	"google.golang.org/grpc"
)

type CustomerService struct {
	client protogen.MainClient
}

func NewCustomerService(client protogen.MainClient) *CustomerService {
	return &CustomerService{client: client}
}

var Customer *CustomerService

func Init() {
	customerServiceConn, err := grpc.Dial(helpers.CustomerServiceURL, grpc.WithInsecure()) // or use helpers.CustomerServiceURL
	if err != nil {
		log.Fatalf("Failed to connect to customer-service: %v", err)
	}

	Customer = NewCustomerService(protogen.NewMainClient(customerServiceConn))
}
