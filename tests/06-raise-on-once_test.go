package tests

import (
	"fmt"
	"testing"
	"time"

	events "github.com/codemodify/systemkit-events"
)

func Test_Raise_On_Once(t *testing.T) {

	const pingEvent = "PING"

	// Emit PING every second
	go func() {
		for {
			time.Sleep(1 * time.Second)
			events.Events().Raise(pingEvent)
		}
	}()

	// Set event HANDLER-1
	const handler1ID = "HANDLER-1"
	const handler1Once = true
	events.Events().On(pingEvent, func() {
		fmt.Println(fmt.Sprintf("[%s] event handled from %s", pingEvent, handler1ID))
	}, handler1Once)

	// Set event HANDLER-2
	const handler2ID = "HANDLER-2"
	const handler2Once = false
	events.Events().On(pingEvent, func() {
		fmt.Println(fmt.Sprintf("[%s] event handled from %s", pingEvent, handler2ID))
	}, handler2Once)

	// Stop the test after 10 seconds
	time.Sleep(10 * time.Second)
}
