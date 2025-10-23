package main

import (
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// handleRequest is your function handler method
func handleRequest(event []byte, ctx context.RuntimeContext) (string, error) {
	logger := ctx.GetLogger()
	logger.Logf("Function invoked with event: %v", event)

	return fmt.Sprintf("Hello, %v!", string(event)), nil
}

// main function starts the runtime with your handler
func main() {
	runtime.Register(handleRequest)
}
