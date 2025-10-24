package main

import (
	"bytes"
	"encoding/json"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/go-api/context"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime"
)

// Example for API Gateway (Dedicated Gateway) handler
func DebugTest(payload []byte, ctx context.RuntimeContext) (interface{}, error) {

	ctx.GetLogger().Logf("GetRequestID: %s", ctx.GetRequestID())
	ctx.GetLogger().Logf("GetRemainingTimeInMilliSeconds: %d", ctx.GetRemainingTimeInMilliSeconds())
	ctx.GetLogger().Logf("GetAccessKey: %s", ctx.GetAccessKey())
	ctx.GetLogger().Logf("GetSecretKey: %s", ctx.GetSecretKey())
	ctx.GetLogger().Logf("GetSecurityAccessKey: %s", ctx.GetSecurityAccessKey())
	ctx.GetLogger().Logf("GetSecuritySecretKey: %s", ctx.GetSecuritySecretKey())
	ctx.GetLogger().Logf("GetFunctionName: %s", ctx.GetFunctionName())
	// ctx.GetLogger().Logf("Alias: %s",ctx.GetUserData(string))
	ctx.GetLogger().Logf("GetRunningTimeInSeconds: %d", ctx.GetRunningTimeInSeconds())
	ctx.GetLogger().Logf("GetVersion: %s", ctx.GetVersion())
	ctx.GetLogger().Logf("GetMemorySize: %d", ctx.GetMemorySize())
	ctx.GetLogger().Logf("GetCPUNumber: %d", ctx.GetCPUNumber())
	ctx.GetLogger().Logf("GetProjectID: %s", ctx.GetProjectID())
	ctx.GetLogger().Logf("GetPackage: %s", ctx.GetPackage())
	ctx.GetLogger().Logf("GetToken: %s", ctx.GetToken())
	ctx.GetLogger().Logf("GetSecurityToken: %s", ctx.GetSecurityToken())
	ctx.GetLogger().Logf("GetAlias: %s", ctx.GetAlias())

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, payload, "", "\t")

	ctx.GetLogger().Logf("Payload: %s", string(prettyJSON.String()))

	return "ok", nil
}

func main() {
	runtime.Register(DebugTest)
}
