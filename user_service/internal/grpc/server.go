package grpc

import (
	"fmt"
	"net"

	pb "github.com/MiracleCanCode/trello_protos/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
	usecase pb.UserServiceServer
}

func New(usecase pb.UserServiceServer) *Server {
	return &Server{
		usecase: usecase,
	}
}

func (s *Server) Conn(addr string) (string, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return "", fmt.Errorf("Conn: failed create tcp listener on %s: %w", addr, err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, s.usecase)

	if err := grpcServer.Serve(lis); err != nil {
		return "", fmt.Errorf("Conn: error work server: %w", err)
	}

	successMessage := fmt.Sprintf("Server start work on %s", addr)
	return successMessage, nil
}
