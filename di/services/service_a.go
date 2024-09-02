package services

import (
	"errors"
	"fmt"
)

type IServiceB interface {
	DoSomethingElse()
}

type ServiceA struct {
	serviceB IServiceB
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (s *ServiceA) SetServiceB(serviceB IServiceB) {
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
