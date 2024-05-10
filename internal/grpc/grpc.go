package grpc

import (
	"log"
	"net"

	"github.com/kubespect/server/internal/grpc/xdp"
	pb "github.com/kubespect/server/protobuf/xdp"
	"google.golang.org/grpc"
)

type Grpc struct {
	grpcServer *grpc.Server
}

func NewGrpc() *Grpc {
	grpcServer := grpc.NewServer()
	pb.RegisterXDPServer(grpcServer, xdp.NewServer())
	s := &Grpc{grpcServer: grpcServer}
	return s
}

func (s *Grpc) Start() {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("grpc server listening on port 9090")
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
