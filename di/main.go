package main

import (
	"di/services"
	"log"
)

func main() {
	serviceA := services.NewServiceA()
	serviceB := services.NewServiceB()

	// Register the services
	serviceA.Register(serviceB)
	serviceB.Register(serviceA)

	// Validate the services
	if err := serviceA.Validate(); err != nil {
		log.Fatalf("Validation failed for ServiceA: %v", err)
	}
	if err := serviceB.Validate(); err != nil {
		log.Fatalf("Validation failed for ServiceB: %v", err)
	}

	// Call service methods
	serviceA.DoSomething()
	serviceB.DoSomethingElse()
}
