package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin framework.
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/events/apig"
)

// invokeIndex is a handler function that processes requests
// to the POST /index route.
func invokeIndex(c *gin.Context) {
	// Read the request body.
	reqBody, _ := c.GetRawData()

	fmt.Printf("Received request body: %s\n", string(reqBody))

	var apigEvent apig.APIGTriggerEvent
	err := json.Unmarshal(reqBody, &apigEvent)
	if err != nil {
		fmt.Println("Unmarshal failed")
		c.String(http.StatusBadRequest, "invalid data")
		return
	}

	fmt.Printf("RequestId:%s\n", apigEvent.RequestContext.RequestId)

	// fill API Gateway response structure.
	returnBody := "Hello from FunctionGraph!"

	c.String(http.StatusOK, returnBody)
}
