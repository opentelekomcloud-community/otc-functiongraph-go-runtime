package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-testutils/context"
)

func Test_APIGTest(t *testing.T) {

	filename := "../resources/apig_base64_event.json"
	event, _ := os.ReadFile(filename)

	fmt.Println("Testing APIG Event Handler: Test_APIGTest")
	t.Log(string(event))

	ctx := context.NewTestContext()

	t.Log(ctx.GetRequestID())

	ret, err := ApigTest(event, ctx)
	t.Log(ret, err)
}

func Test_APIGTest_PathParams(t *testing.T) {

	filename := "../resources/apig_event_with_params.json"
	event, _ := os.ReadFile(filename)

	fmt.Println("Testing APIG Event Handler: Test_APIGTest_PathParams")
	t.Log(string(event))

	ctx := context.NewTestContext()

	t.Log(ctx.GetRequestID())

	ret, err := ApigTest(event, ctx)
	t.Log(ret, err)
}
