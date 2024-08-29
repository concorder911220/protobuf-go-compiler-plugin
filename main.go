package main

import (
	"circular_dependency/common/a"
	"circular_dependency/common/b"
	"circular_dependency/common/eventbus"
	"fmt"
)

func main() {
	eventBus := eventbus.NewEventBus()
	bService := b.NewB(eventBus)
	aService := a.NewA(bService, eventBus)

	aService.CreateRoom([]string{"Hello World"})
	roomInfo := aService.GetRoomInfo("authorizedRoom")
	fmt.Println(roomInfo)

	bService.CreatePost("authorizedRoom", []string{"New Post"})
}
