package main

import (
	"bytes"
	"container-http/src/fghttpmiddleware"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin" // Import the Gin framework
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// go init function is called before main function.
func init() {
	log.Println("init in main.go ")
}

// Logger is a middleware function used
// to record request and response information.
func LogRequestResponse(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, logger := fghttpmiddleware.FGContext(c)
		logger.Infow("FunctionName", "functionName", ctx.GetFuncName())

		start := time.Now()

		reqBody, _ := c.GetRawData()

		logger.Infow("Request", "method", c.Request.Method, "uri", c.Request.RequestURI, "body", string(reqBody))

		// Assign Back the request body as body can only be read once
		// by default in Gin framework.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		respBody := string(rune(c.Writer.Size()))

		logger.Infow("Response", "method", c.Request.Method, "uri", c.Request.RequestURI, "body", respBody, "latency", latency)
	}
}

func main() {
	// Create a Gin router.
	router := gin.New()

	// Configure zap logger to output logs in JSON
	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "ts"

	// config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.00000Z"))
	})

	// Create a logger
	logger, _ := config.Build()
	defer logger.Sync()

	// Create a sugared logger for easier logging with structured context.
	sugar := logger.Sugar()

	// Add middleware for FunctionGraph context and logging
	router.Use(fghttpmiddleware.InitFGMiddleware(sugar))

	// LogRequestResponse middleware logs the request and response information.
	router.Use(LogRequestResponse(sugar))

	// Recovery middleware recovers from any panics and writes
	// a 500 if there was one.
	router.Use(gin.Recovery())

	// Registers a route for handling HTTP POST requests at the path `/index`.
	// When a client sends a
	// POST request to `/index`, the Gin framework will call
	// the `InvokeIndexPOST` function to handle the request.
	router.POST("/index", InvokeIndexPOST)

	router.GET("/index", InvokeIndexGET)

	err := router.Run(":8000") // Starts the HTTP server, listening on port 8000.
	if err != nil {
		return
	}
}
