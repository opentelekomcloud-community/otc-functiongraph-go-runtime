package apig

import (
	"encoding/base64"
)

type APIGTriggerResponse struct {
	// Message body
	Body string `json:"body"`

	// The final returned Http response header
	Headers map[string]string `json:"headers"`

	// Http status code, int type
	StatusCode int `json:"statusCode"`

	// Whether the body is base64 encoded, bool type
	IsBase64Encoded bool `json:"isBase64Encoded"`
}

func (r *APIGTriggerResponse) SetBase64EncodedBody(body string) {
	r.Body = base64.StdEncoding.EncodeToString([]byte(body))
}
