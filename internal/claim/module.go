package claim

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/app"
	"github.com/NewGlassbiller/gb-services-insurance/internal/claim/infra/repo"
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
	grpcgen.RegisterclaimServer(s, m.service)
}

func (m *Module) RegisterGRPCGateway(s *gbgateway.Server) {
	s.RegisterHandler(grpcgen.RegisterclaimHandlerFromEndpoint)
}

func (m *Module) RegisterCLI(app gbcli.AppInterface) {
	app.AddCommand(
		&cli.Command{
			Name:  "claim",
			Usage: "Operations on claim",
			Subcommands: []*cli.Command{
				{
					Name:  "load",
					Usage: "Load claim from shared database",
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
