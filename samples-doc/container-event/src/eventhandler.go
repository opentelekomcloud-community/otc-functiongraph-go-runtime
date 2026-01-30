package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin" // Import the Gin framework.
)

// go init function is called before main function.
func init() {
	fmt.Println("init in main.go ")
}

// Logger is a middleware function used
// to record request and response information.
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		logTimeFormat := "2006-01-02T15:04:05.000Z"
		start := time.Now()

		// Get the request ID from the request header.
		requestID := c.GetHeader("X-Cff-Request-Id")
		reqBody, _ := c.GetRawData()

		fmt.Printf("%s [INFO] Request:  %s %s %s %s\n", time.Now().UTC().Format(logTimeFormat), requestID, c.Request.Method,
			c.Request.RequestURI, reqBody)

		// Assign Back the request body as body can only be read once
		// by default in Gin framework.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		respBody := string(rune(c.Writer.Size()))
		fmt.Printf("%s [INFO] Response: %s %s %s %s (%v)\n", time.Now().UTC().Format(logTimeFormat), requestID, c.Request.Method,
			c.Request.RequestURI, respBody, latency)
	}
}

// main is the entry point of the application.
// In Production set Environment variable
// GIN_MODE=release
func main() {

	// Create a Gin router.
	router := gin.New()

	// Global middleware
	router.Use(Logger())

	// Recovery middleware recovers from any panics and writes
	// a 500 if there was one.
	router.Use(gin.Recovery())

	// Register a route (/invoke) that processes HTTP POST requests.
	// When FunctionGraph sends a POST request to /invoke, the Gin framework
	// calls the invokeSampleData function to process the request.
	router.POST("/invoke", invokeSampleData)

	// Handle undefined routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND" + c.Request.Method + " " + c.Request.RequestURI, "message": "Page not found"})
	})

	// Start the HTTP server and listen to port 8000.
	err := router.Run(":8000")
	if err != nil {
		return
	}
}
