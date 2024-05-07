package xdp

import (
	"encoding/binary"
	pb "github.com/kubespect/server/protobuf/xdp"
	"io"
	"log"
	"net"
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
		if err == io.EOF {
			return stream.SendAndClose(&pb.XDPResponse{
				Interface: "ens18",
			})
		}
		if err != nil {
			log.Printf("Error receiving packet: %v", err)
		}
		//log.Printf("Received packet %d: %v", packetCount, packet)
		log.Printf("%s:%d -> %s:%d, seq: %d, ack: %d, flags: %d, window: %d", PrintIP(packet.SrcIP), packet.SrcPort, PrintIP(packet.DstIP), packet.DstPort, packet.Seq, packet.Ack, packet.Flags, packet.Window)
	}
}

func PrintIP(ip uint32) string {
	result := make(net.IP, 4)
	binary.BigEndian.PutUint32(result, ip)

	return result.String()
}