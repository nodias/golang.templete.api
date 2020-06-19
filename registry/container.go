package registry

import (
	"github.com/nodias/golang.templete.api/domain/service"
	"github.com/nodias/golang.templete.api/interface/persistence/postgres"
	"github.com/nodias/golang.templete.api/usecase"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if builder.Add([]di.Def{
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	repo := postgres.NewUserRepository()
	service := service.NewUserService(repo)
	return usecase.NewUserUsecase(repo, service), nil
}
