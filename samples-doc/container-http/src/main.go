package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin" // Import the Gin framework
)

// go init function is called before main function.
func init() {
	fmt.Println("init in main.go ")
}

// Logger is a middleware function used to log request and response information.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqBody, _ := c.GetRawData()
		fmt.Printf("[INFO] Request: %s %s %s\n", c.Request.Method, c.Request.RequestURI, reqBody)

		// Assign Back the request body as body can only be read once
		// by default in Gin framework.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		respBody := string(rune(c.Writer.Size()))
		fmt.Printf("[INFO] Response: %s %s %s (%v)\n", c.Request.Method, c.Request.RequestURI, respBody, latency)
	}
}

func main() {
	router := gin.Default() // Creates a default Gin router

	router.Use(Logger()) // Use the Logger middleware

	router.POST("/index", invokeIndex) // Registers a route for handling HTTP POST requests at the path `/index`. When a client sends a POST request to `/index`, the Gin framework will call the `index` function to handle the request.

	err := router.Run(":8000") // Starts the HTTP server, listening on port 8000.
	if err != nil {
		return
	}
}
