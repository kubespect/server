package xdp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/kubespect/protobuf/xdp"
)

type Server struct {
	pb.UnimplementedXDPServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) XDPStream(stream pb.XDP_XDPStreamServer) error {
	for {
		packet, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			if err := stream.SendAndClose(&pb.XDPResponse{
				Interface: "ens18",
			}); err != nil {
				return fmt.Errorf("Error sending response: %w", err)
			}
			return nil
		}
		if err != nil {
			return fmt.Errorf("Error receiving packet: %w", err)
		}
		if packet != nil {
			log.Printf("%s:%d -> %s:%d, seq: %d, ack: %d, flags: %d, window: %d", printIP(packet.SrcIP), packet.SrcPort, printIP(packet.DstIP), packet.DstPort, packet.Seq, packet.Ack, packet.Flags, packet.Window)
		}
	}
}

func printIP(ip uint32) string {
	result := make(net.IP, 4)
	binary.BigEndian.PutUint32(result, ip)

	return result.String()
}
