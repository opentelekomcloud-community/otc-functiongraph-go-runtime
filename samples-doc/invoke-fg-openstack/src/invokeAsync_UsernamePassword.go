package invoke_fg_openstack

import (
	"encoding/json"
	"fmt"
	"os"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/invoke"
)

func InvokeAsync_UsernamePassword() {
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
		Region: "eu-de",
	})
	if err != nil {
		fmt.Printf("Failed to create FunctionGraph client: %v", err)
		panic(err)
	}

	body := map[string]string{
		"key": "Hello World of OTC",
	}

	funcURN := fmt.Sprintf("urn:fss:eu-de:%s:function:default:DefaultPython3_10:latest", provider.ProjectID)
	fmt.Println("Function URN: " + funcURN)

	jsonString, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Failed to marshal body: %v", err)
		panic(err)
	}
	fmt.Println("Body: " + string(jsonString))

	result, err := invoke.LaunchAsync(client, funcURN, body)
	if err != nil {
		fmt.Printf("Failed to invoke function: %v", err)
		panic(err)
	}

	fmt.Printf("Function RequestID: %s\n", result.RequestID)

}
