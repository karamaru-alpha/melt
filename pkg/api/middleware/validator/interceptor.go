package validator

import (
	"context"

	"google.golang.org/grpc"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type validator interface {
	Validate() error
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, merrors.Wrap(err, merrors.InvalidArgument)
			}
		}
		return handler(ctx, req)
	}
}
