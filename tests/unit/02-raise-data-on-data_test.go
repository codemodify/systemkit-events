package tests

import (
	"fmt"
	"testing"
	"time"

	events "github.com/codemodify/systemkit-events"
)

func Test_RaiseData_OnData(t *testing.T) {

	const pingEvent = "PING"
	const pingData = "PING-DATA"

	// Emit PING every second
	go func() {
		for {
			time.Sleep(1 * time.Second)
			events.Events().EmitWithData(pingEvent, []byte(pingData))
		}
	}()

	// Set event HANDLER-1
	const handler1ID = "HANDLER-1"
	events.Events().OnWithData(pingEvent, func(data []byte) {
		fmt.Println(fmt.Sprintf("[%s] event with data [%s] handled from %s", pingEvent, string(data), handler1ID))
	})

	// Set event HANDLER-2
	const handler2ID = "HANDLER-2"
	events.Events().OnWithData(pingEvent, func(data []byte) {
		fmt.Println(fmt.Sprintf("[%s] event with data [%s] handled from %s", pingEvent, string(data), handler2ID))
	})

	// Stop the test after 10 seconds
	time.Sleep(10 * time.Second)
}
