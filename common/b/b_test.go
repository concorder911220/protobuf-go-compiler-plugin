package b

import (
	"circular_dependency/common/eventbus"
	"testing"
)

func TestCreatePost(t *testing.T) {
	eb := eventbus.NewEventBus()
	bService := NewB(eb)

	// Subscribe to checkUser event
	eb.Subscribe("checkUser", func(data interface{}) interface{} {
		return true // Simulate authorized user
	})

	// Test post creation
	bService.CreatePost("authorizedRoom", []string{"New Post"}) // Should succeed

	// Test unauthorized post creation
	eb.Subscribe("checkUser", func(data interface{}) interface{} {
		return false // Simulate unauthorized user
	})
	bService.CreatePost("unauthorizedRoom", []string{"New Post"}) // Should fail
}
