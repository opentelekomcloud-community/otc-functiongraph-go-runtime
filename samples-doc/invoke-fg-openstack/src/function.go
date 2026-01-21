package invoke_fg_openstack

import (
	"encoding/base64"
	"fmt"
	"os"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/fgs/v2/function"
)

func CreateFunction_UsernamePassword() {

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

	createOpts := function.CreateOpts{
		Name:       functionName,
		Package:    functionApp,
		Runtime:    "Python3.10",
		Timeout:    200,
		Handler:    "index.handler",
		MemorySize: 128,
		CodeType:   "inline",
		FuncCode: &function.FuncCode{
			File: base64.StdEncoding.EncodeToString([]byte(appCode)),
		},
	}

	createResp, err := function.Create(client, createOpts)
	if err != nil {
		fmt.Printf("Failed to create function: %v", err)
		panic(err)
	}

	fmt.Printf("Function created: %+v\n", createResp)
}

func DeleteFunction_UsernamePassword() {

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

	projectId := provider.ProjectID

	// no version in deletion!
	funcURN := fmt.Sprintf("urn:fss:%s:%s:function:%s:%s", region, projectId, functionApp, functionName)

	err = function.Delete(client, funcURN)
	if err != nil {
		fmt.Printf("Failed to delete function: %v", err)
		panic(err)
	}

	fmt.Println("Function deleted successfully")

}
