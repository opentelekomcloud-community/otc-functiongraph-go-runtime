package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin" // Import the Gin framework.
)

// SampleData represents the structure of the expected JSON payload:
//
//	{ "key": "value"}
//
// in the request body.
type SampleData struct {
	Key string `json:"key"`
}

// invokeSampleData is a function that processes
// POST requests sent to the /invoke route.
func invokeSampleData(c *gin.Context) {
	// Read the request body.
	reqBody, _ := c.GetRawData()

	fmt.Printf("Received request body: %s\n", string(reqBody))

	// Print all request headers.
	// for name, values := range c.Request.Header {
	// 	// Loop over all values for the name.
	// 	for _, value := range values {
	// 		fmt.Printf(" %s: %s\n", name, value)
	// 	}
	// }

	var sampleData SampleData

	err := json.Unmarshal(reqBody, &sampleData)
	if err != nil {
		fmt.Println("Unmarshal failed")
		c.String(http.StatusBadRequest, "invalid data")
		return
	}

	fmt.Printf("Key: %s\n", sampleData.Key)

	c.String(http.StatusOK, "Received key: %s", sampleData.Key)

}
