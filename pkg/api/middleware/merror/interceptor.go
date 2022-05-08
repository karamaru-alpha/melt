package merror

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/karamaru-alpha/melt/pkg/merrors"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		res, err := handler(ctx, req)
		if err == nil {
			return res, nil
		}

		pattern := merrors.Unknown
		message := err.Error()
		var meltError *merrors.MeltError
		if errors.As(err, &meltError) {
			pattern = meltError.ErrorPattern
			message = meltError.Message()
		}
		return res, status.Error(pattern.GRPCStatusCode, message)
	}
}
