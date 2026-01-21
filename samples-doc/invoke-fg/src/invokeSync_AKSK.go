package invoke_fg

//
// This sample shows how to invoke a FunctionGraph function using AK/SK authentication.
// It demonstrates how to set up the request, sign it with AK/SK, and handle the response.
//
// Before running the sample, ensure that you have set the following environment variables:
// - OTC_SDK_PROJECTID: Your OpenTelekomCloud Project ID
// - OTC_TENANT_NAME: Your OpenTelekomCloud region (tenant name)
// - OTC_SDK_AK: Your Access Key
// - OTC_SDK_SK: Your Secret Key
//

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core"
)

var httpClientAKSK = &http.Client{}

// calls FunctionGraph function using AK/SK authentication
func InvokeSync_AKSK() error {
	fmt.Println("Calling FunctionGraph function with Access Key and Secret Key...")

	projectID := os.Getenv("OTC_SDK_PROJECTID")
	region := os.Getenv("OTC_TENANT_NAME")
	ak := os.Getenv("OTC_SDK_AK")
	sk := os.Getenv("OTC_SDK_SK")

	// FunctionGraph endpoint
	fgEndpoint := fmt.Sprintf("https://functiongraph.%s.otc.t-systems.com", region)

	// Function name/version and application
	functionName := "DefaultPython3_10"
	functionVersion := "latest"
	functionApp := "default"

	functionURN := fmt.Sprintf("urn:fss:%s:%s:function:%s:%s:%s", region, projectID, functionApp, functionName, functionVersion)
	invokeURI := fmt.Sprintf("%s/v2/%s/fgs/functions/%s/invocations", fgEndpoint, projectID, functionURN)

	// X-Cff-Log-Type:
	// "tail": 4KB logs will be returned
	// "": no logs will be returned
	xCffLogType := "tail"

	// X-CFF-Request-Version:
	// "v0" response body in text format
	// "v1" response body in json format
	xCffRequestVersion := "v1"

	// Set request body
	body := map[string]interface{}{
		"key": "Hello World of OTC",
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error marshaling request body: %s", err)
		panic(err)
	}

	// Create request
	req, err := http.NewRequest("POST", invokeURI, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("error creating request: %s", err)
		panic(err)
	}

	// Add required headers
	req.Header.Set("X-Project-Id", projectID)
	req.Header.Set("Content-Type", "application/json;charset=utf8")
	req.Header.Set("X-Cff-Log-Type", xCffLogType)
	req.Header.Set("X-CFF-Request-Version", xCffRequestVersion)
	req.Header.Set("Host", req.URL.Host)

	// create signer
	// see: https://docs.otc.t-systems.com/developer/api_guidelines/calling_apis/ak_sk_authentication/index.html#ak-sk-authentication
	signer := core.Signer{
		Key:    ak,
		Secret: sk,
	}
	// sign the request
	signer.Sign(req)

	fmt.Printf("Invoking FunctionGraph function using URI %s\n", invokeURI)

	// Send the request
	resp, err := httpClientAKSK.Do(req)
	if err != nil {
		fmt.Printf("request error: %s", err)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response code %d, %s\n", resp.StatusCode, resp.Status)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response: %s", err)
		panic(err)
	}

	// Handle response based on X-CFF-Request-Version
	if xCffRequestVersion == "v1" {
		// OutputType is json
		var data map[string]interface{}
		if err := json.Unmarshal(responseBody, &data); err != nil {
			fmt.Printf("error parsing response: %s", err)
			panic(err)
		}

		fmt.Println("---- Result ----")
		if result, ok := data["result"]; ok {
			resultJSON, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(resultJSON))
		}

		if xCffLogType == "tail" {
			fmt.Println("---- Logs ----")
			if log, ok := data["log"]; ok {
				fmt.Println(log)
			}
		}
	} else {
		// OutputType is text
		fmt.Println(string(responseBody))
	}

	return nil
}
