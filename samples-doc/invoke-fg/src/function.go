package invoke_fg

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// create a FunctionGraph using API calls
func CreateFunction_UsernamePassword() error {
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

	// FunctionGraph endpoint
	fgEndpoint := fmt.Sprintf("https://functiongraph.%s.otc.t-systems.com", region)

	createURI := fmt.Sprintf("%s/v2/%s/fgs/functions", fgEndpoint, projectID)
	// Set request body
	body := map[string]interface{}{
		"func_name":   functionName,
		"package":     functionApp,
		"runtime":     "Python3.10",
		"timeout":     200,
		"handler":     "index.handler",
		"memory_size": 128,
		"code_type":   "inline",
		"func_code": map[string]string{
			"file": base64.StdEncoding.EncodeToString([]byte(appCode)),
		},
	}

	requestBody, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("error marshaling request body: %s", err)
		panic(err)
	}

	req, err := http.NewRequest("POST", createURI, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("error creating request: %s", err)
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AUTH-Token", token)

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

	var data map[string]interface{}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	fmt.Println("---- Result ----")
	resultJSON, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(resultJSON))

	return nil
}

// delete a FunctionGraph using API calls
func DeleteFunction_UsernamePassword() {
	region := os.Getenv("OTC_SDK_REGION")
	userName := os.Getenv("OTC_USER_NAME")
	userPassword := os.Getenv("OTC_USER_PASSWORD")
	domainName := os.Getenv("OTC_DOMAIN_NAME")
	projectID := os.Getenv("OTC_SDK_PROJECTID")
	authURL := os.Getenv("OTC_IAM_ENDPOINT")

	var httpClient = &http.Client{}

	token, err := getTokenUserNamePassword(*httpClient, userName, userPassword, domainName, authURL, projectID)
	if err != nil {
		fmt.Printf("failed to get token: %s", err)
		panic(err)
	}

	// FunctionGraph endpoint
	fgEndpoint := fmt.Sprintf("https://functiongraph.%s.otc.t-systems.com", region)

	functionURN := fmt.Sprintf("urn:fss:%s:%s:function:%s:%s1", region, projectID, functionApp, functionName)
	deleteURI := fmt.Sprintf("%s/v2/%s/fgs/functions/%s", fgEndpoint, projectID, functionURN)

	req, err := http.NewRequest("DELETE", deleteURI, nil)
	if err != nil {
		fmt.Printf("error creating request: %s", err)
		panic(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-AUTH-Token", token)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("request error: %s", err)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response code %d, %s\n", resp.StatusCode, resp.Status)

}
