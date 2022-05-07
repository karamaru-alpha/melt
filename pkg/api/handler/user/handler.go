package user

import (
	"context"

	"github.com/karamaru-alpha/melt/pkg/api/usecase/user"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type handler struct {
	userInteractor user.Interactor
}

func New(userInteractor user.Interactor) pb.UserServer {
	return &handler{
		userInteractor: userInteractor,
	}
}

func (h *handler) Create(ctx context.Context, req *pb.UserCreateRequest) (*pb.UserCreateResponse, error) {
	token, err := h.userInteractor.Create(ctx, req.GetName())
	if err != nil {
		return nil, merrors.Stack(err)
	}
	return &pb.UserCreateResponse{Token: token}, nil
}
