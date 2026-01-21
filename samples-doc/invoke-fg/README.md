# Samples on how to call FunctionGraph function using Go

## Prerequisites

1. Needed environment variables set

   | Name              | Description |
   | ----------------- | ----------- |
   | OTC_SDK_PROJECTID | Project ID
   | OTC_SDK_REGION    | Region, e.g. "eu-de"
   | OTC_SDK_AK        | Access Key
   | OTC_SDK_SK        | Secret Key
   | OTC_USER_NAME     | User name
   | OTC_USER_PASSWORD | User password
   | OTC_DOMAIN_NAME   | Domain name
   | OTC_IAM_ENDPOINT  | IAM Endpoint, e.g. https://iam.eu-de.otc.t-systems.com/v3

2. Python FunctionGraph event-function created and deployed with:
   * Runtime: "Python 3.10"
   * Name: "DefaultPython3_10"
   * Handler: index.handler
   * Application: default
  
## Call Functiongraph using Username and Password

See src/invokeUsernamePassword.go


## Call Functiongraph using AK/SK

See src/invokeAKSK.go

