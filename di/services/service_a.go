package services

import (
	"di/common"
	"errors"
	"fmt"
)

type ServiceA struct {
	serviceB common.IServiceB
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (s *ServiceA) SetServiceB(serviceB common.IServiceB) {
	s.serviceB = serviceB
}

func (s *ServiceA) Validate() error {
	if s.serviceB == nil {
		return errors.New("ServiceA: ServiceB is not set")
	}
	return nil
}

func (s *ServiceA) DoSomething() {

	fmt.Println("ServiceA is doing something.")

}

func (s *ServiceA) Register(serviceB common.IServiceB) {
	s.SetServiceB(serviceB)
}
