package registry

import (
	"google.golang.org/grpc"

	"github.com/karamaru-alpha/melt/pkg/api/registry/auth"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
)

func Register(s *grpc.Server) {
	pb.RegisterAuthServer(s, auth.DI())
}
