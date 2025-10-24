package main

import (
	"encoding/json"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/timer"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

func handleRequest(event []byte, ctx context.RuntimeContext) (interface{}, error) {
	var timerEvent timer.TimerTriggerEvent

	logger := ctx.GetLogger()

	err := json.Unmarshal(event, &timerEvent)
	if err != nil {
		logger.Logf("Unmarshal of event %s failed: %v", timerEvent.String(), err)
		return "invalid data", err
	}

	logger.Logf("Function invoked with event: %v", timerEvent.String())
	return timerEvent.String(), nil
}

// main function starts the runtime with your handler
func main() {
	runtime.Register(handleRequest)
}
