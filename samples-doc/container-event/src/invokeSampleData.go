package main

import (
	"container-event-sample/src/fgeventmiddleware"
	"encoding/json"
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
	ctx, logger := fgeventmiddleware.FGContext(c)

	// Read the request body.
	reqBody, _ := c.GetRawData()

	logger.Infow("Received request body", "body", string(reqBody))

	var sampleData SampleData

	err := json.Unmarshal(reqBody, &sampleData)
	if err != nil {

		logger.Errorw("Unmarshal failed", "error", err)
		c.String(http.StatusBadRequest, "invalid data")
		return
	}

	logger.Infow("RequestID", "requestID", ctx.GetRequestID())

	logger.Infow("Processing request", "key", sampleData.Key)

	c.String(http.StatusOK, "Received key: %s", sampleData.Key)
}
