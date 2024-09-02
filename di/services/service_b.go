package services

import (
	"di/common"
	"errors"
	"fmt"
)

type ServiceB struct {
	serviceA common.IServiceA
}

func NewServiceB() *ServiceB {
	return &ServiceB{}
}
func (s *ServiceB) SetServiceA(serviceA common.IServiceA) {
	s.serviceA = serviceA
}

func (s *ServiceB) Validate() error {
	if s.serviceA == nil {
		return errors.New("ServiceB: ServiceA is not set")
	}
	return nil
}

func (s *ServiceB) DoSomethingElse() {
	fmt.Println("ServiceB is doing something else.")
}

func (s *ServiceB) Register(serviceA common.IServiceA) {
	s.SetServiceA(serviceA)
}
