package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-testutils/context"
)

func Test_CtsTest(t *testing.T) {

	filename := "../resources/cts_event.json"
	event, _ := os.ReadFile(filename)

	fmt.Println("Testing CTS Event Handler")
	t.Log(string(event))

	ctx := context.NewTestContext()

	t.Log(ctx.GetRequestID())

	ret, err := CtsTest(event, ctx)
	t.Log(ret, err)
}
