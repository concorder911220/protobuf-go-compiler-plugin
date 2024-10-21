package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NewGlassbiller/gb-services-insurance/internal/user/app"
	"github.com/NewGlassbiller/gb-services-insurance/internal/user/infra/repo"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Module struct {
	service *Service
	command *app.Command
}

var instance *Module

func NewModule(service *Service, command *app.Command) *Module {
	return &Module{
		service: service,
		command: command,
	}
}

func Singleton() *Module {
	if instance == nil {
		module, err := wireModule()
		if err != nil {
			panic(err)
		}
		instance = module
	}
	return instance
}

func wireModule() (*Module, error) {
	wire.Build(
		repo.New,
		app.NewQuery,
		app.NewCommand,
		NewService,
		NewModule,
	)

	return &Module{}, nil
}

func (m *Module) RegisterGRPC(s *gbgrpc.Server) {
	grpcgen.RegisteruserServer(s, m.service)
}

func (m *Module) RegisterGRPCGateway(s *gbgateway.Server) {
	s.RegisterHandler(grpcgen.RegisteruserHandlerFromEndpoint)
}

func (m *Module) RegisterCLI(app gbcli.AppInterface) {
	app.AddCommand(
		&cli.Command{
			Name:  "user",
			Usage: "Operations on user",
			Subcommands: []*cli.Command{
				{
					Name:  "load",
					Usage: "Load user from shared database",
					Action: func(c *cli.Context) error {
						res, err := m.command.Load()
						// Encode response to JSON
						responseJson, err := json.MarshalIndent(res, "", "  ")
						if err != nil {
							return err
						}
						fmt.Println(string(responseJson))

						return nil
					},
				},
			},
		},
	)
}
