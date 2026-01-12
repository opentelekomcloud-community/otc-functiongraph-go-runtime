package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/timer"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Timer handler
func TimerTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var timerEvent timer.TimerTriggerEvent
	err := json.Unmarshal(payload, &timerEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	return timerEvent.String(), nil
}

func main() {
	runtime.Register(TimerTest)
}
