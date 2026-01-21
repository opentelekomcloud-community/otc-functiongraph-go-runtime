package invoke_fg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// calls FunctionGraph function using username and password authentication
func InvokeAsync_UsernamePassword() error {
	userName := os.Getenv("OTC_USER_NAME")
	userPassword := os.Getenv("OTC_USER_PASSWORD")
	domainName := os.Getenv("OTC_DOMAIN_NAME")
	projectID := os.Getenv("OTC_SDK_PROJECTID")
	region := os.Getenv("OTC_SDK_REGION")
	authURL := os.Getenv("OTC_IAM_ENDPOINT")

	token, err := getTokenUserNamePassword(userName, userPassword, domainName, authURL, projectID)
	if err != nil {
		fmt.Printf("failed to get token: %s", err)
		panic(err)
	}

	fmt.Println("Obtained Token:", token)

	fgEndpoint := fmt.Sprintf("https://functiongraph.%s.otc.t-systems.com", region)

	functionApp := "default"
	functionName := "DefaultPython3_10"
	functionVersion := "latest"

	functionURN := fmt.Sprintf("urn:fss:%s:%s:function:%s:%s:%s", region, projectID, functionApp, functionName, functionVersion)
	invokeURI := fmt.Sprintf("%s/v2/%s/fgs/functions/%s/invocations-async", fgEndpoint, projectID, functionURN)

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

	// Send the request
	resp, err := httpClientUserNamePassword.Do(req)
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

	var data map[string]interface{}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	fmt.Println("---- Result ----")
	if result, ok := data["result"]; ok {
		resultJSON, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(resultJSON))
	}

	return nil
}
