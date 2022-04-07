package registry

import (
	"google.golang.org/grpc"

	"github.com/karamaru-alpha/melt/pkg/api/registry/user"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
)

func Register(s *grpc.Server) {
	pb.RegisterUserServer(s, user.DI())
}
