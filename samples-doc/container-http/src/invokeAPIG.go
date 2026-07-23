package main

import (
	"container-http/src/fghttpmiddleware"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin framework.
)

// invokeIndexPOST is a handler function that processes requests
// to the POST /index route.
func InvokeIndexPOST(c *gin.Context) {
	ctx, logger := fghttpmiddleware.FGContext(c)

	// Read the request body.
	reqBody, _ := c.GetRawData()

	logger.Infof("Received request [%s] body: %s\n", ctx.GetRequestID(), string(reqBody))

	c.Request.Header.Get("")

	// fill API Gateway response structure.
	returnBody := "Hello from FunctionGraph!"

	c.String(http.StatusOK, returnBody)
}

// invokeIndexGET is a handler function that processes requests
// to the GET /index route.
func InvokeIndexGET(c *gin.Context) {
	// Read query parameters
	name := c.Query("name")
	if name == "" {
		name = "World"
	}

	returnBody := fmt.Sprintf("Hello, %s! This is a GET request.", name)

	c.String(http.StatusOK, returnBody)
}
