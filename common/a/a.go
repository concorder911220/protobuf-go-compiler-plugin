package a

import (
	"circular_dependency/common/b"
	"circular_dependency/common/eventbus"
)

type A struct {
	bService *b.B
	eventBus *eventbus.EventBus
}

func NewA(bService *b.B, eventBus *eventbus.EventBus) *A {
	a := &A{bService: bService, eventBus: eventBus}
	eventBus.Subscribe("checkUser", a.checkUser)
	return a
}

func (a *A) checkUser(data interface{}) interface{} {
	roomID := data.(string)
	isAuthorized := roomID == "authorizedRoom" // Example logic
	return isAuthorized
}

func (a *A) GetRoomInfo(roomID string) interface{} {
	posts := a.bService.GetPosts()
	return map[string]interface{}{
		"roomInfo": nil,
		"posts":    posts,
	}
}

func (a *A) CreateRoom(posts []string) {
	a.bService.CreatePost("authorizedRoom", posts)
}
