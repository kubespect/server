package xdp

import (
	"io"
	"log"

	pb "github.com/kubespect/server/protobuf/xdp"
)

type Server struct {
	pb.UnimplementedXDPServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) XdpStream(stream pb.XDP_XDPStreamServer) error {
	var packetCount uint32
	for {
		packet, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.XDPResponse{
				Interface: "ens18",
			})
		}
		if err != nil {
			log.Printf("Error receiving packet: %v", err)
		}
		packetCount++
		log.Printf("Received packet %d: %v", packetCount, packet)
	}
}
