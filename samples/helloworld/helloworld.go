package main

import (
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

func main() {
	runtime.Register(Handler)
}

func Handler(payload []byte, ctx context.RuntimeContext) (interface{}, error) {

	ctx.GetLogger().Logf("Received payload: %s", string(payload))

	return "Hello, OTC FunctionGraph!", nil
}
