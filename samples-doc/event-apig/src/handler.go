package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-events/apig"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for API Gateway (Dedicated Gateway) handler
func ApigTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {
	var apigEvent apig.APIGTriggerEvent
	err := json.Unmarshal(payload, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		return "invalid data", err
	}
	ctx.GetLogger().Logf("payload:%s", apigEvent.String())

	// body processing example

	var body string

	if apigEvent.IsBase64Encoded {
		body = apigEvent.GetRawBody()
		ctx.GetLogger().Logf("Decoded body: %s", body)
	} else {
		body = apigEvent.Body
		ctx.GetLogger().Logf("Body: %s", body)
	}

	// path parameters example
	pathParameters := apigEvent.PathParameters
	if pathParameters != nil {
		for key, value := range pathParameters {
			ctx.GetLogger().Logf("Path parameter - %s: %s", key, value)
		}
	}

	// headers example
	headers := apigEvent.Headers
	if headers != nil {
		for key, value := range headers {
			ctx.GetLogger().Logf("Header - %s: %s", key, value)
		}
	}

	var returnBody string

	if apigEvent.IsBase64Encoded {
		returnBodyBytes, err := base64.StdEncoding.DecodeString(body)
		if err != nil {
			return nil, err
		}
		returnBody = string(returnBodyBytes)
	} else {
		returnBody = body
	}

	apigResp := apig.APIGTriggerResponse{
		Body: returnBody,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		StatusCode: 200,
	}

	return apigResp, nil
}

func main() {
	runtime.Register(ApigTest)
}
