package main

import (
	"encoding/json"
	"fmt"

	obs "github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/obss3"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for OBS Service handler
func ObsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var obsEvent obs.OBSS3TriggerEvent
	err := json.Unmarshal(payload, &obsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", obsEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(ObsTest)
}
