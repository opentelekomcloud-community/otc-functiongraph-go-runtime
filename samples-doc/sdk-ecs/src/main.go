package main

/**
SDK ECS example
This example shows how to use the OpenTelekomCloud API SDK to start, stop or reboot an ECS instance using the FunctionGraph service.

For API reference, see:
- https://docs.otc.t-systems.com/elastic-cloud-server/api-ref/apis_recommended/batch_operations/restarting_ecss_in_a_batch.html#restarting-ecss-in-a-batch

The function reads the following user data parameters:
- ECS_ENDPOINT_URL: The endpoint URL of the ECS service (optional, default value is used if not provided)
- ECS_INSTANCE_ID: The ID of the ECS instance to be managed
- ECS_ACTION: The action to be performed on the instance (start, stop, reboot)
- ECS_ACTION_TYPE: The type of action for stop and reboot (SOFT, HARD) (not required for start action)

The function uses the security credentials (temporary AK/SK/Token) from the runtime context to sign the API requests.

*/

import (
	"bytes"
	"encoding/json"

	"io"
	"strings"

	"errors"
	"net/http"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"

	"github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core"
)

func handlerECS(payload []byte, ctx context.RuntimeContext) (interface{}, error) {

	// get ProjectID from context
	project_id := ctx.GetProjectID()

	// get ECS endpoint from user data or use default value
	var endpoint = ctx.GetUserData("ECS_ENDPOINT_URL")
	if endpoint == "" {
		endpoint = "https://ecs.eu-de.otc.t-systems.com"
	}

	// get InstanceID from user data
	instance_id := ctx.GetUserData("ECS_INSTANCE_ID")
	if (instance_id) == "" {
		ctx.GetLogger().Logf("ECS_INSTANCE_ID not defined")
		return "ECS_INSTANCE_ID not defined", errors.New("ECS_INSTANCE_ID not defined")
	}

	action := strings.ToLower(ctx.GetUserData("ECS_ACTION"))
	if action == "" {
		ctx.GetLogger().Logf("ECS_ACTION not defined, supported: start, stop, reboot")
		return "ECS_ACTION not defined", errors.New("ECS_ACTION not defined")
	}
	if action != "start" && action != "stop" && action != "reboot" {
		ctx.GetLogger().Logf("ECS_ACTION not supported, supported: start, stop, reboot")
		return "ECS_ACTION not supported", errors.New("ECS_ACTION not supported")
	}

	actionType := strings.ToUpper(ctx.GetUserData("ECS_ACTION_TYPE"))
	if actionType == "" && action != "start" {
		ctx.GetLogger().Logf("ECS_ACTION_TYPE not defined, supported: SOFT, HARD")
		return "ECS_ACTION_TYPE not defined", errors.New("ECS_ACTION_TYPE not defined")
	}

	if actionType != "SOFT" && actionType != "HARD" && action != "start" {
		ctx.GetLogger().Logf("ECS_ACTION_TYPE not supported, supported: SOFT, HARD")
		return "ECS_ACTION_TYPE not supported", errors.New("ECS_ACTION_TYPE not supported")
	}

	var jsonStr map[string]interface{}
	//
	switch action {
	case "start":
		jsonStr = map[string]interface{}{
			"os-start": map[string]interface{}{
				"servers": []map[string]string{
					{"id": instance_id},
				},
			},
		}

	case "stop":
		jsonStr = map[string]interface{}{
			"os-stop": map[string]interface{}{
				"type": actionType,
				"servers": []map[string]string{
					{"id": instance_id},
				},
			},
		}

	case "reboot":
		jsonStr = map[string]interface{}{
			"reboot": map[string]interface{}{
				"type": actionType,
				"servers": []map[string]string{
					{"id": instance_id},
				},
			},
		}
	}

	// Marshal the JSON body
	jsonData, err := json.Marshal(jsonStr)
	if err != nil {
		ctx.GetLogger().Logf("JSON marshal failed:", err)
		return "invalid json", err
	}

	// Create the HTTP request
	request, err := http.NewRequest("POST", endpoint+"/v1/"+project_id+"/cloudservers/action",
		bytes.NewReader(jsonData))

	if err != nil {
		ctx.GetLogger().Logf("ECS request failed:", err)
		return "invalid request", err
	}

	// set headers
	request.Header.Set("content-type", "application/json; charset=utf-8")

	if project_id != "" {
		// To access resources in a sub-project (e.g. eu_de/myproject)
		// by calling APIs, X-Project-Id of "eu_de/myproject" is needed
		request.Header.Set("X-Project-Id", project_id)
	}

	// create signer
	signer := core.Signer{
		Key:           ctx.GetSecurityAccessKey(),
		Secret:        ctx.GetSecuritySecretKey(),
		SecurityToken: ctx.GetSecurityToken(),
	}
	// sign the request
	signer.Sign(request)

	// create http client
	client := http.DefaultClient

	// Send the request
	resp, err := client.Do(request)
	if err != nil {
		ctx.GetLogger().Logf("ECS request failed:", err)
		return "invalid request", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// prepare return value
	ret := map[string]interface{}{
		"statusCode": resp.StatusCode,
		"body":       string(body),
	}

	return ret, nil
}

func main() {
	runtime.Register(handlerECS)
}
