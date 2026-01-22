package invoke_fg

// Before running the sample, ensure that you have set the following environment variables:
// - OTC_USER_NAME: Your OpenTelekomCloud username
// - OTC_USER_PASSWORD: Your OpenTelekomCloud user password
// - OTC_DOMAIN_NAME: Your OpenTelekomCloud domain name
// - OTC_SDK_PROJECTID: Your OpenTelekomCloud Project ID
// - OTC_SDK_REGION: Your OpenTelekomCloud region (e.g., "eu-de")
// - OTC_IAM_ENDPOINT: Your OpenTelekomCloud IAM endpoint URL

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// calls FunctionGraph function using username and password authentication
func InvokeSync_UsernamePassword() error {
	userName := os.Getenv("OTC_USER_NAME")
	userPassword := os.Getenv("OTC_USER_PASSWORD")
	domainName := os.Getenv("OTC_DOMAIN_NAME")
	projectID := os.Getenv("OTC_SDK_PROJECTID")
	region := os.Getenv("OTC_SDK_REGION")
	authURL := os.Getenv("OTC_IAM_ENDPOINT")

	var httpClient = &http.Client{}

	token, err := getTokenUserNamePassword(*httpClient, userName, userPassword, domainName, authURL, projectID)
	if err != nil {
		fmt.Printf("failed to get token: %s", err)
		panic(err)
	}

	fmt.Println("Obtained Token:", token)

	fgEndpoint := fmt.Sprintf("https://functiongraph.%s.otc.t-systems.com", region)

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

	req, err := http.NewRequest("POST", invokeURI, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("error creating request: %s", err)
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AUTH-Token", token)
	req.Header.Set("X-Cff-Log-Type", xCffLogType)
	req.Header.Set("X-CFF-Request-Version", xCffRequestVersion)

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("request error: %s", err)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response code %d, %s\n", resp.StatusCode, resp.Status)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if xCffRequestVersion == "v1" {
		// OutputType is json
		var data map[string]interface{}
		if err := json.Unmarshal(responseBody, &data); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
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
