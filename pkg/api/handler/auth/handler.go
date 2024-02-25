package auth

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/karamaru-alpha/melt/pkg/api/usecase/auth"

	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
	"github.com/karamaru-alpha/melt/pkg/domain/proto/model"
	"github.com/karamaru-alpha/melt/pkg/mcontext"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

type handler struct {
	authInteractor auth.Interactor
}

func New(authInteractor auth.Interactor) pb.AuthServer {
	return &handler{
		authInteractor: authInteractor,
	}
}

func (h *handler) Signup(ctx context.Context, req *pb.AuthSignupRequest) (*pb.AuthSignupResponse, error) {
	accessToken, refreshToken, err := h.authInteractor.Signup(ctx, req.GetName())
	if err != nil {
		return nil, merrors.Stack(err)
	}
	return &pb.AuthSignupResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (h *handler) RefreshToken(ctx context.Context, req *pb.AuthRefreshTokenRequest) (*pb.AuthRefreshTokenResponse, error) {
	accessToken, err := h.authInteractor.RefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, merrors.Stack(err)
	}
	return &pb.AuthRefreshTokenResponse{AccessToken: accessToken}, nil
}

func (h *handler) Get(ctx context.Context, _ *emptypb.Empty) (*model.User, error) {
	mctx := mcontext.Extract(ctx)
	return &model.User{
		Id: mctx.GetUserID(),
	}, nil
}
