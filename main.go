package main

import (
	"fmt"

	"github.com/javonlee22/event-handler/events"
)

func test1(data interface{}) {
	fmt.Println("Printing from test1")
}

func test2(data interface{}) {
	fmt.Println("Printing from test2")
}

func main() {
	eventManager := events.GetEventManager()
	eventManager.Subscribe("test", "test1", test1)
	eventManager.Subscribe("test", "test2", test2)
	eventManager.Emit("test", "hello")
}
