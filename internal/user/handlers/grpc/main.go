package grpc

import (
	"log"

	"github.com/pusrenk/auth-service/pkg/helpers"
	"google.golang.org/grpc"
)

var (
	CustomerService *CustomerService
)

func Init() {
	customerServiceConn, err := grpc.Dial(helpers.CustomerServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}

	CustomerService = CustomerService(customerServiceConn)
}
