package main

import (
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/dds"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for Document Database Service handler
func DdsTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var ddsEvent dds.DDSTriggerEvent
	err := json.Unmarshal(payload, &ddsEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", ddsEvent.String())
	return "ok", nil
}

func main() {
	runtime.Register(DdsTest)
}
