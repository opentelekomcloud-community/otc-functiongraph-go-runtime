/*
API Gateway (APIG) is an API hosting service that helps enterprises to build, manage,
and deploy APIs at any scale. With APIG, your function can be invoked through HTTPS by
using a custom REST API and a specified backend. You can map each API operation (such as, GET and PUT)
to a specific function. APIG invokes the relevant function when an HTTPS request is sent to the API backend.

For more information about how to use HTTPS calls to trigger functions, see
https://docs.otc.t-systems.com/function-graph/umn/creating_triggers/using_an_apig_dedicated_trigger.html
*/
package apig

import (
	"encoding/base64"
	"fmt"
)

type APIGTriggerEvent struct {
	// is body base64 encoded? Default value: true
	IsBase64Encoded bool `json:"isBase64Encoded"`

	// Http request method
	HttpMethod string `json:"httpMethod"`

	// Http request path
	Path string `json:"path"`

	// Http request body
	Body string `json:"body"`

	// Path parameters
	PathParameters map[string]string `json:"pathParameters"`

	// Request information, including the API gateway configuration, request ID,
	// authentication information, and source.
	RequestContext APIGRequestContext `json:"requestContext"`

	// Http requestheaders
	Headers map[string]string `json:"headers"`

	// Query strings configured in APIG and their actual values
	QueryStringParameters map[string]string `json:"queryStringParameters"`

	// Userdata set in APIG custom authentication
	UserData string `json:"user_data"`
}

func (e *APIGTriggerEvent) GetRawBody() string {
	decoded, err := base64.StdEncoding.DecodeString(e.Body)
	if err != nil {
		return ""
	}
	return string(decoded)
}

func (e *APIGTriggerEvent) String() string {
	return fmt.Sprintf(`APIGTriggerEvent{
                                 isBase64Encoded=%v,
                                 httpMethod='%v',
                                 path='%v',
                                 body='%v',
                                 pathParamters=%v,
                                 requestContext=%v,
                                 headers=%v,
                                 queryStringParameters=%v,
                                 user_data=%+v,
                              }`, e.IsBase64Encoded, e.HttpMethod, e.Path, e.Body, e.PathParameters, e.RequestContext,
		e.Headers, e.QueryStringParameters, e.UserData)
}
