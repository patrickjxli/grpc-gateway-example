package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Stoakes/grpc-gateway-example/echopb"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

// MicroServer represents a microservice server instance
type MicroServer struct {
	serverName string
	lis        net.Listener
	httpServer *http.Server
	grpcServer *grpc.Server
}

// New returns a microserver instance
func New(serverName string, l net.Listener) *MicroServer {
	return &MicroServer{
		serverName: serverName,
		lis:        l,
	}
}

// Echo implementation
// Implement EchoService interface
func (m *MicroServer) Echo(c context.Context, s *pb.EchoMessage) (*pb.EchoMessage, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Value)
	return s, nil
}

func newServer() *MicroServer {
	return new(MicroServer)
}

// -----------------------------------------------------------------------------

// Start the microserver
func (ms *MicroServer) Start() error {
	var err error

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// tcpMuxer
	tcpMux := cmux.New(ms.lis)

	// Connection dispatcher rules
	grpcL := tcpMux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := tcpMux.Match(cmux.HTTP1Fast())

	// initialize gRPC server instance
	ms.grpcServer, err = prepareGRPC(ctx)
	if err != nil {
		log.Fatalln("Unable to initialize gRPC server instance")
		return err
	}

	// initialize HTTP server
	ms.httpServer, err = prepareHTTP(ctx, ms.serverName)
	if err != nil {
		log.Fatalln("Unable to initialize HTTP server instance")
		return err
	}

	// Start servers
	go func() {
		if err := ms.grpcServer.Serve(grpcL); err != nil {
			log.Fatalln("Unable to start external gRPC server")
		}
	}()
	go func() {
		if err := ms.httpServer.Serve(httpL); err != nil {
			log.Fatalln("Unable to start HTTP server")
		}
	}()

	return tcpMux.Serve()
}
