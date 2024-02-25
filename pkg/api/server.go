package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/karamaru-alpha/melt/pkg/api/middleware/auth"
	"github.com/karamaru-alpha/melt/pkg/api/middleware/merror"
	"github.com/karamaru-alpha/melt/pkg/api/middleware/recovery"
	"github.com/karamaru-alpha/melt/pkg/api/middleware/validator"
	"github.com/karamaru-alpha/melt/pkg/api/registry"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type Config struct {
	Port string
}

type Server struct {
	config *Config
	server *grpc.Server
}

func Serve(c *Config) *Server {
	server := NewServer(c)
	if server == nil {
		return nil
	}
	server.ServeAndWait()
	return server
}

func NewServer(c *Config) *Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				auth.UnaryServerInterceptor(),
				merror.UnaryServerInterceptor(),
				recovery.UnaryServerInterceptor(),
				validator.UnaryServerInterceptor(),
			),
		),
	)
	reflection.Register(server)
	registry.Register(server)
	return &Server{
		config: c,
		server: server,
	}
}

func (s *Server) ServeAndWait() {
	ctx := context.Background()
	lis, err := s.listen()
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		s.serve(lis)
		cancel()
	}()
	defer log.Println("Shutdown...")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ctx.Done():
	case sig := <-signalCh:
		log.Printf("Received signal. %v\n", sig)
		s.server.GracefulStop()
	}
}

func (s *Server) listen() (net.Listener, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		return nil, merrors.Stack(err)
	}
	return lis, nil
}

func (s *Server) serve(lis net.Listener) {
	log.Printf("grpc server started on :%s", s.config.Port)
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}
