package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/smn"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Simple Notification Service handler
func SmnTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var smnEvent smn.SMNTriggerEvent
	err := json.Unmarshal(payload, &smnEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", smnEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(SmnTest)
}
