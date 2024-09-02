package services

import (
	"errors"
	"fmt"
)

type IServiceA interface {
	DoSomething()
}

type ServiceB struct {
	serviceA IServiceA
}

func NewServiceB() *ServiceB {
	return &ServiceB{}
}

func (s *ServiceB) SetServiceA(serviceA IServiceA) {
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
