package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/lts"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Log Tank Service handler
func LtsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ltsEvent lts.LTSTriggerEvent
	err := json.Unmarshal(payload, &ltsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ltsEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(LtsTest)
}
