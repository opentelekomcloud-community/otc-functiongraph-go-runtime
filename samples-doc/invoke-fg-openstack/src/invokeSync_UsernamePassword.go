package invoke_fg_openstack

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/invoke"
)

func InvokeSync_UsernamePassword() {

	region := os.Getenv("OTC_SDK_REGION")

	opts := golangsdk.AuthOptions{
		IdentityEndpoint: "https://iam.eu-de.otc.t-systems.com/v3",
		Username:         os.Getenv("OTC_USER_NAME"),
		Password:         os.Getenv("OTC_USER_PASSWORD"),
		DomainName:       os.Getenv("OTC_DOMAIN_NAME"),
		TenantID:         os.Getenv("OTC_SDK_PROJECTID"),
	}

	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		fmt.Printf("Failed to get client: %v", err)
		panic(err)
	}

	client, err := openstack.NewFuncGraph(provider, golangsdk.EndpointOpts{
		Region: region,
	})
	if err != nil {
		fmt.Printf("Failed to create FunctionGraph client: %v", err)
		panic(err)
	}

	body := map[string]interface{}{
		"key":  "Hello World of OTC",
		"test": map[string]int{"a": 1, "b": 2},
	}

	funcURN := fmt.Sprintf("urn:fss:%s:%s:function:default:DefaultPython3_10:latest", region, provider.ProjectID)
	fmt.Println("Function URN: " + funcURN)

	jsonString, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Failed to marshal body: %v", err)
		panic(err)
	}
	fmt.Println("Body: " + string(jsonString))

	launchSyncHeaders := invoke.LaunchSyncHeaders{
		LogType: "tail",
	}

	result, resultHeaders, err := invoke.LaunchSync(client, funcURN, body, launchSyncHeaders)
	if err != nil {
		fmt.Printf("Failed to invoke function: %v", err)
		panic(err)
	}

	fmt.Printf("Function RequestID: %s\n", result.RequestID)
	fmt.Printf("Function Result: %s\n", result.Result)
	fmt.Printf("Function Log: %s\n", result.Log)
	fmt.Printf("Function Status: %d\n", result.Status)

	fmt.Println("RequestId:", resultHeaders.RequestId)

	fmt.Printf("Result HEADERS: %+v", resultHeaders)

	fmt.Println("InvokeSummary:", resultHeaders.InvokeSummary)

	h := resultHeaders.AdditionalHeaders
	for key := range h {
		fmt.Printf("AdditionalHeaders: %s %s\n", key, h[key][0])
	}

	decodedLog, err := base64.StdEncoding.DecodeString(resultHeaders.FunctionLog)
	if err != nil {
		fmt.Printf("Failed to decode function log: %v", err)
		panic(err)
	}
	fmt.Println(string(decodedLog))

	result2, err := invoke.LaunchAsync(client, funcURN, body)
	if err != nil {
		fmt.Printf("Failed to invoke function asynchronously: %v", err)
		panic(err)
	}

	fmt.Printf("result2: %s\n", result2.RequestID)

}
