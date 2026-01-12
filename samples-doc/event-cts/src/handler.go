package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/cts"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Cloud Trace Service handler
func CtsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ctsEvent cts.CTSTriggerEvent
	err := json.Unmarshal(payload, &ctsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ctsEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(CtsTest)
}
