package b

import (
	"circular_dependency/common/eventbus"
	"fmt"
)

type B struct {
	eventBus *eventbus.EventBus
}

func NewB(eventBus *eventbus.EventBus) *B {
	return &B{eventBus: eventBus}
}

func (b *B) CreatePost(roomID string, posts []string) {
	isAuthorized := b.eventBus.Publish("checkUser", roomID)

	if isAuthorized == nil || !isAuthorized.(bool) {
		fmt.Println("User is not authorized to create a post in this room.")
		return
	}

	fmt.Println("Post created successfully:", posts)
}

func (b *B) GetPosts() []string {
	return []string{"Post 1", "Post 2"}
}
