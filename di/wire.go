//go:build wireinject
// +build wireinject

package main

import (
	"di/services"

	"github.com/google/wire"
)

type ServiceContainer struct {
	ServiceA *services.ServiceA
	ServiceB *services.ServiceB
}

func InitializeServices() (*ServiceContainer, error) {
	wire.Build(
		NewServiceContainer,
	)
	return &ServiceContainer{}, nil
}

func NewServiceContainer() (*ServiceContainer, error) {
	serviceA := services.NewServiceA()
	serviceB := services.NewServiceB()

	// Inject each other
	serviceA.SetServiceB(serviceB)
	serviceB.SetServiceA(serviceA)

	return &ServiceContainer{
		ServiceA: serviceA,
		ServiceB: serviceB,
	}, nil
}
