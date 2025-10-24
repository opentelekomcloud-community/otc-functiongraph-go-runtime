package apig

import (
	"fmt"
)

type APIGRequestContext struct {
	// API ID
	ApiId string `json:"apiId"`

	// The requestId of this API request
	RequestId string `json:"requestId"`

	// Release environment name
	Stage string `json:"stage"`

	// Source IP address
	SourceIp string `json:"sourceIp"`
}

func (rc APIGRequestContext) String() string {
	return fmt.Sprintf(`APIGRequestContext{
                                 apiId='%s',
                                 requestId='%s',
                                 stage='%s',
                                 sourceIp='%s',
                               }`, rc.ApiId, rc.RequestId, rc.Stage, rc.SourceIp)
}
