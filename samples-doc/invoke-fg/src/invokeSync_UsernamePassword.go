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

var httpClientUserNamePassword = &http.Client{}

// getTokenUserNamePassword retrieves an authentication token using username and password
func getTokenUserNamePassword(userName, userPassword, domainName, authURL, projectID string) (string, error) {
	tokenURI := authURL + "/auth/tokens?v3/auth/tokens?nocatalog=true"

	// Build auth request body
	authBody := map[string]interface{}{
		"auth": map[string]interface{}{
			"identity": map[string]interface{}{
				"methods": []string{"password"},
				"password": map[string]interface{}{
					"user": map[string]interface{}{
						"name":     userName,
						"password": userPassword,
						"domain": map[string]interface{}{
							"name": domainName,
						},
					},
				},
			},
			"scope": map[string]interface{}{
				"domain": map[string]interface{}{
					"name": domainName,
				},
				"project": map[string]interface{}{
					"id": projectID,
				},
			},
		},
	}

	requestBody, err := json.Marshal(authBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling auth body: %w", err)
	}

	req, err := http.NewRequest("POST", tokenURI, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClientUserNamePassword.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting token: %w", err)
	}
	defer resp.Body.Close()

	token := resp.Header.Get("X-Subject-Token")

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		responseText, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error getting token: %s\n", string(responseText))
		return token, fmt.Errorf("token request failed with status: %d", resp.StatusCode)
	}

	return token, nil
}

// calls FunctionGraph function using username and password authentication
func InvokeSync_UsernamePassword() error {
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
