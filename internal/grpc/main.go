package grpc

import (
	"log"
	"sync"

	"github.com/pusrenk/auth-service/internal/protobuf/customer-service/protogen"
	rpc "github.com/pusrenk/auth-service/internal/user/handlers/grpc/customer-service"
	"github.com/pusrenk/auth-service/pkg/helpers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn     *grpc.ClientConn
	once     sync.Once
	Customer *rpc.CustomerService
)

func Init() {
	once.Do(func() {
		var err error
		conn, err = grpc.NewClient(helpers.CustomerServiceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to customer-service: %v", err)
		}

		Customer = rpc.NewCustomerService(protogen.NewCustomerRpcClient(conn))
	})
}
