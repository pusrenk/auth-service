package server

import (
	"context"
	"log"
	"net"

	"github.com/pusrenk/auth-service/configs"
	"github.com/pusrenk/auth-service/database"
	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/internal/user/handlers"
	"github.com/pusrenk/auth-service/internal/user/repositories"
	"github.com/pusrenk/auth-service/internal/user/services"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	config     *configs.Config
}

func NewServer(cfg *configs.Config) *Server {
	return &Server{
		grpcServer: grpc.NewServer(),
		config:     cfg,
	}
}

func (s *Server) SetupDependencies() error {
	// Initialize Redis
	redisClient, err := database.InitRedis(s.config)
	if err != nil {
		return err
	}

	// Initialize dependencies
	userRedisRepo := repositories.NewUserRedisRepository(redisClient)
	userService := services.NewUserService(userRedisRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Register the service
	protogen.RegisterMainServer(s.grpcServer, userHandler)

	return nil
}

func (s *Server) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	log.Printf("gRPC server starting on port %s...", port)
	return s.grpcServer.Serve(listener)
}

func (s *Server) Stop(ctx context.Context) {
	log.Println("Shutting down gRPC server...")
	s.grpcServer.GracefulStop()
}
