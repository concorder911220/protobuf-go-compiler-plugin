package main

import (
	"di/services"
	"log"
)

func main() {
	serviceA := services.NewServiceA()
	serviceB := services.NewServiceB()

	serviceA.SetServiceB(serviceB)
	serviceB.SetServiceA(serviceA)

	if err := serviceA.Validate(); err != nil {
		log.Fatalf("Validation failed for ServiceA: %v", err)
	}
	if err := serviceB.Validate(); err != nil {
		log.Fatalf("Validation failed for ServiceB: %v", err)
	}

	serviceA.DoSomething()
	serviceB.DoSomethingElse()
}
