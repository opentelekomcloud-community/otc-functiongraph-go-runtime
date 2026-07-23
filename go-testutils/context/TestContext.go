package context

// TestContext.go provides a test implementation of the RuntimeContext interface
// for use in unit tests.
// It allows developers to simulate the function execution context
// without needing to deploy the function to the actual runtime environment.

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime/common"
)

// TestFunctionLogger implements common.RuntimeLogger interface
type TestFunctionLogger struct {
	RequestID string
}

func (l *TestFunctionLogger) Log(message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000000000")
	fmt.Printf("[%s] %s %s\n", timestamp, l.RequestID, message)
}

func (l *TestFunctionLogger) Logf(format string, args ...interface{}) {
	l.Log(fmt.Sprintf(format, args...))
}

// TestContext implements context.RuntimeContext interface for testing
type TestContext struct {
	UserData             map[string]string
	FuncStart            time.Time
	RunningTimeInSeconds int
	RequestID            string
	FunctionName         string
	Version              string
	MemorySize           int
	CPUNumber            int
	ProjectID            string
	Package              string
	Alias                string
	Alias2               string
	AccessKey            string
	SecretKey            string
	SecurityAccessKey    string
	SecuritySecretKey    string
	Token                string
	SecurityToken        string
	StreamRunID          string
	StreamEnable         string
	StreamToken          string
	StreamAddr           string
	OriginVersionTag     string
	WorkflowStateID      string
	WorkflowID           string
	WorkflowRunID        string
}

func (ctx TestContext) GetRequestID() string { return ctx.RequestID }

func (ctx TestContext) GetRemainingTimeInMilliSeconds() int {
	elapsed := time.Since(ctx.FuncStart).Milliseconds()
	remaining := (ctx.RunningTimeInSeconds * 1000) - int(elapsed)
	return remaining
}

func (ctx TestContext) GetAccessKey() string         { return ctx.AccessKey }
func (ctx TestContext) GetSecretKey() string         { return ctx.SecretKey }
func (ctx TestContext) GetSecurityAccessKey() string { return ctx.SecurityAccessKey }
func (ctx TestContext) GetSecuritySecretKey() string { return ctx.SecuritySecretKey }
func (ctx TestContext) GetFunctionName() string      { return ctx.FunctionName }

func (ctx TestContext) GetUserData(key string) string {
	if val, ok := ctx.UserData[key]; ok {
		return val
	}
	return ""
}

func (ctx TestContext) GetLogger() common.RuntimeLogger {
	return &TestFunctionLogger{}
}

func (ctx TestContext) GetRunningTimeInSeconds() int { return ctx.RunningTimeInSeconds }
func (ctx TestContext) GetVersion() string           { return ctx.Version }
func (ctx TestContext) GetMemorySize() int           { return ctx.MemorySize }
func (ctx TestContext) GetCPUNumber() int            { return ctx.CPUNumber }
func (ctx TestContext) GetProjectID() string         { return ctx.ProjectID }
func (ctx TestContext) GetPackage() string           { return ctx.Package }
func (ctx TestContext) GetToken() string             { return ctx.Token }
func (ctx TestContext) GetSecurityToken() string     { return ctx.SecurityToken }
func (ctx TestContext) GetAlias() string             { return ctx.Alias }
func (ctx TestContext) GetStreamRunID() string       { return ctx.StreamRunID }
func (ctx TestContext) GetStreamEnable() string      { return ctx.StreamEnable }
func (ctx TestContext) GetStreamToken() string       { return ctx.StreamToken }
func (ctx TestContext) GetStreamAddr() string        { return ctx.StreamAddr }
func (ctx TestContext) GetOriginVersionTag() string  { return ctx.OriginVersionTag }
func (ctx TestContext) GetWorkflowStateID() string   { return ctx.WorkflowStateID }
func (ctx TestContext) GetWorkflowID() string        { return ctx.WorkflowID }
func (ctx TestContext) GetWorkflowRunID() string     { return ctx.WorkflowRunID }

// NewTestContext creates a new TestContext with default/mock values
func NewTestContext() TestContext {
	return TestContext{
		AccessKey:            "mock-access-key",
		Alias:                "mock-alias",
		CPUNumber:            1,
		FuncStart:            time.Now(),
		FunctionName:         "mock-function-name",
		MemorySize:           128,
		OriginVersionTag:     "mock-origin-version-tag",
		Package:              "default",
		ProjectID:            "mock-project-id",
		RequestID:            uuid.New().String(),
		RunningTimeInSeconds: 30,
		SecretKey:            "mock-secret-key",
		SecurityAccessKey:    "mock-security-access-key",
		SecuritySecretKey:    "mock-security-secret-key",
		SecurityToken:        "mock-security-token",
		StreamAddr:           "mock-stream-addr",
		StreamEnable:         "true",
		StreamRunID:          "mock-stream-run-id",
		StreamToken:          "mock-stream-token",
		Token:                "mock-token",
		Version:              "latest",
		WorkflowID:           "mock-workflow-id",
		WorkflowRunID:        "mock-workflow-run-id",
		WorkflowStateID:      "mock-workflow-state-id",

		UserData: map[string]string{
			"key1": "value1",
		},
	}
}
