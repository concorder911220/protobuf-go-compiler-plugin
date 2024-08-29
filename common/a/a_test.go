package a

import (
	"circular_dependency/common/b"
	"circular_dependency/common/eventbus"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	eb := eventbus.NewEventBus()
	bService := b.NewB(eb)
	aService := NewA(bService, eb)

	// Subscribe to checkUser event
	eb.Subscribe("checkUser", func(data interface{}) interface{} {
		return true // Simulate authorized user
	})

	aService.CreateRoom([]string{"Hello World"})
}

func TestCheckUser(t *testing.T) {
	eb := eventbus.NewEventBus()
	bService := b.NewB(eb)
	aService := NewA(bService, eb)

	// Test authorization logic
	isAuthorized := aService.checkUser("authorizedRoom")
	if isAuthorized != true {
		t.Error("Expected user to be authorized for 'authorizedRoom'")
	}

	isAuthorized = aService.checkUser("unauthorizedRoom")
	if isAuthorized != false {
		t.Error("Expected user to be unauthorized for 'unauthorizedRoom'")
	}
}
