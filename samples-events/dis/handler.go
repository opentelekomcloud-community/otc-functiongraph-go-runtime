package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/dis"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Data Ingestion Service handler
func DisTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var disEvent dis.DISTriggerEvent
	err := json.Unmarshal(payload, &disEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", disEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(DisTest)
}
