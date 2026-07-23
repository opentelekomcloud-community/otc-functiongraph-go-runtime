package fghttpmiddleware

import (
	"github.com/gin-gonic/gin" // Import the Gin framework.
	context "github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-http/httpcontext"
	"go.uber.org/zap"
)

const (
	loggerKey  = "fg.LoggerWithRequestID"
	contextKey = "fg.FGContext"
)

// FGContext retrieves the FunctionGraph context and logger from the Gin context.
func FGContext(c *gin.Context) (context.ContextProvider, *zap.SugaredLogger) {
	fgCtx, ok := c.Get(contextKey)

	if !ok {
		panic("FGContext not found in context! Middleware '" + contextKey + "' might not be applied.")
	}

	logger, ok := c.Get(loggerKey)

	if !ok {
		panic("Logger not found in context! Middleware '" + loggerKey + "' might not be applied.")
	}

	return fgCtx.(context.ContextProvider), logger.(*zap.SugaredLogger)
}

// InitFGMiddleware is a middleware function to add request ID to the logger
func InitFGMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Initialize the context environment and provider, and set them in the Gin context.
		contextProvider := context.GetContextProvider(c.Request)
		c.Set(contextKey, contextProvider)

		// Add request ID to the logger
		requestID := contextProvider.GetRequestID()
		logger = logger.With("requestID", requestID)

		// Add function version to the logger
		funcVersion := contextProvider.GetFuncVersion()
		logger = logger.With("funcVersion", funcVersion)

		// add logger with fields to the Gin context
		c.Set(loggerKey, logger)

		c.Next()
	}
}
