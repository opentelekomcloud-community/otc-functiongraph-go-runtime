package main

import (
	"os"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

func main() {
	runtime.Register(Handler)
}

func Handler(payload []byte, ctx context.RuntimeContext) (interface{}, error) {

	ctx.GetLogger().Logf("Received payload: %s", string(payload))

	for _, e := range os.Environ() {
		// pair := strings.SplitN(e, "=", 2)
		ctx.GetLogger().Logf(e)
	}

	return "Hello, OTC FunctionGraph!", nil
}
