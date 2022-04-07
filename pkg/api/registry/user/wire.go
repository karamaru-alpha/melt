//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"

	handler "github.com/karamaru-alpha/melt/pkg/api/handler/user"
	interactor "github.com/karamaru-alpha/melt/pkg/api/usecase/user"
	pb "github.com/karamaru-alpha/melt/pkg/domain/proto/api"
	service "github.com/karamaru-alpha/melt/pkg/domain/service/user"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql"
	"github.com/karamaru-alpha/melt/pkg/infra/mysql/repository"
	"github.com/karamaru-alpha/melt/pkg/util"
)

func DI() pb.UserServer {
	wire.Build(
		handler.New,
		interactor.New,
		service.New,
		repository.NewUserRepository,
		mysql.New,
		util.NewUILDGenerator,
		mysql.NewDBTxManager,
	)

	return nil
}
