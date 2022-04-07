package api

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/karamaru-alpha/melt/pkg/api/handler/user"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
	"github.com/karamaru-alpha/melt/pkg/logging/app"
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
	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterTestServer(server, user.New())
	return &Server{
		config: c,
		server: server,
	}
}

func (s *Server) ServeAndWait() {
	logger := app.GetLogger()
	ctx := context.Background()
	lis, err := s.listen(ctx)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("failed to listen: %v", err))
		return
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		s.serve(ctx, lis)
		cancel()
	}()
	defer logger.Info(ctx, "Shutdown...")

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-ctx.Done():
	case sig := <-signalCh:
		logger.Info(ctx, fmt.Sprintf("Received signal. %v", sig))
		s.server.GracefulStop()
	}
}

func (s *Server) listen(ctx context.Context) (net.Listener, error) {
	logger := app.GetLogger()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.config.Port))
	if err != nil {
		logger.Error(ctx, err.Error())
		return nil, err
	}
	return lis, nil
}

func (s *Server) serve(ctx context.Context, lis net.Listener) {
	logger := app.GetLogger()
	logger.Info(ctx, fmt.Sprintf("grpc server started on :%s", s.config.Port))
	if err := s.server.Serve(lis); err != nil {
		logger.Error(ctx, "failed to serve")
	}
}
