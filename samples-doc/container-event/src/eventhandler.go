package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin" // Import the Gin framework.
)

// Logger is a middleware function used
// to record request and response information.
func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		logTimeFormat := "2006-01-02T15:04:05.000Z"
		start := time.Now()

		requestID := c.GetHeader("X-Cff-Request-Id")
		reqBody, _ := c.GetRawData()

		fmt.Printf("%s [INFO] Request: %s %s %s %s\n", time.Now().UTC().Format(logTimeFormat), requestID, c.Request.Method,
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

// go init function is called before main function.
func init() {
	fmt.Println("init in main.go ")
}

// main is the entry point of the application.
// In Production set Environment variable
// GIN_MODE=release
func main() {

	// Create a default Gin router.
	// router := gin.Default()
	router := gin.New()
	// Middlewares
	{
		router.Use(gin.Recovery())
		// requestid middleware
		router.Use(
			requestid.New(
				requestid.WithGenerator(func() string {
					return "test"
				}),
				requestid.WithCustomHeaderStrKey("x-cff-request-id"),
			),
		)

		//middleware which enhance Gin request logger to include 'RequestID'
		router.Use(gin.LoggerWithConfig(GetLoggerConfig(nil, nil, nil)))
	}

	// Use the Logger middleware.
	// router.Use(Logger())

	// Register a route (/invoke) that processes HTTP POST requests.
	// When FunctionGraph sends a POST request to /invoke, the Gin framework
	// calls the invoke function to process the request.
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

// GetLoggerConfig return gin.LoggerConfig which will write the logs to specified io.Writer with given gin.LogFormatter.
// By default gin.DefaultWriter = os.Stdout
// reference: https://gin-gonic.com/en/docs/examples/custom-log-format/
func GetLoggerConfig(formatter gin.LogFormatter, output io.Writer, skipPaths []string) gin.LoggerConfig {
	if formatter == nil {
		formatter = GetDefaultLogFormatterWithRequestID()
	}
	return gin.LoggerConfig{
		Formatter: formatter,
		Output:    output,
		SkipPaths: skipPaths,
	}
}

// GetDefaultLogFormatterWithRequestID returns gin.LogFormatter with 'RequestID'
func GetDefaultLogFormatterWithRequestID() gin.LogFormatter {
	return func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN-debug] %s [%s] - [%s] \"%s %s %s %d %s\" %s %s\n",
			param.TimeStamp.Format("2006-01-02T15:04:05.000Z"),
			param.Request.Header.Get("x-cff-request-id"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
}
