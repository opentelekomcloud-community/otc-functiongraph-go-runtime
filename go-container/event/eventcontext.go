package eventcontext

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	debugMode = true
)

func getIntHeader(header http.Header, name string) int {
	result, err := strconv.Atoi(header.Get(name))
	if err != nil {
		log.Println("execute strconv.Atoi failed for" + name + " .")
	}
	return result
}

// contextEnv holds the environment variables related to the function context.
type contextEnv struct {
	rtTZ        string // TZ set in dockerfile
	rtPaasPodID string // PAAS_POD_ID
}

// contextHTTP holds the HTTP header values related to the function context.
type contextHTTP struct {
	region            string // X-Cff-Region  [cn]
	requestID         string // X-Cff-Request-Id
	accessKey         string // X-Cff-Access-Key
	secretKey         string // X-Cff-Secret-Key
	securityAccessKey string // X-Cff-Security-Access-Key
	securitySecretKey string // X-Cff-Security-Secret-Key
	authToken         string // X-Cff-Auth-Token
	fcStartTime       int64
	rtRemainTime      int
	securityToken     string // X-Cff-Security-Token

	originVersionTag string // X-Cff-Origin-Version-Tag [!a1]
	workflowStateId  string // X-Cff-Workflow-State-Id
	workflowId       string // X-Cff-Workflow-Id
	workflowRunId    string // X-Cff-Workflow-Run-Id
	funcName         string // X-Cff-Func-Name [0@default@go_go-doc-sample-container-event]
	packageName      string // X-Cff-Package
	funcVersion      string // X-Cff-Func-Version
	projectID        string // X-Cff-Project-Id
	memory           int    // X-Cff-Memory
	timeout          int    // X-Cff-Timeout
}

// ContextProvider provides access to function context information through getter methods.
type ContextProvider struct {
	ctxEnv      *contextEnv
	ctxHTTPHead *contextHTTP
}

func newContextEnv() *contextEnv {
	return &contextEnv{
		rtTZ:        os.Getenv("TZ"),
		rtPaasPodID: os.Getenv("PAAS_POD_ID"),
	}
}

func getCurrentTime() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetContextProvider creates and returns a ContextProvider from an HTTP request.
// This is the main entry point for obtaining function context information.
func GetContextProvider(req *http.Request) ContextProvider {
	return ContextProvider{
		ctxEnv:      newContextEnv(),
		ctxHTTPHead: newContextHTTP(req),
	}
}

// GetTimeZone returns the time zone of the function execution environment.
// This is typically set in the Dockerfile using the TZ environment variable.
func (ctxProvider ContextProvider) GetTimeZone() string {
	return ctxProvider.ctxEnv.rtTZ
}

// GetPaasPodID returns the PAAS_POD_ID of the function execution environment.
func (ctxProvider ContextProvider) GetPaasPodID() string {
	return ctxProvider.ctxEnv.rtPaasPodID
}

func (ctxProvider ContextProvider) GetRemainingTimeInMilliSeconds() int {
	currentTime := getCurrentTime()
	usedTime := int(currentTime - ctxProvider.ctxHTTPHead.fcStartTime)
	timeout := ctxProvider.ctxHTTPHead.timeout * 1000
	if usedTime < timeout {
		ctxProvider.ctxHTTPHead.rtRemainTime = timeout - usedTime
	} else {
		ctxProvider.ctxHTTPHead.rtRemainTime = 0
	}
	return ctxProvider.ctxHTTPHead.rtRemainTime
}

func (ctxProvider ContextProvider) GetRegion() string {
	return ctxProvider.ctxHTTPHead.region
}

func (ctxProvider ContextProvider) GetRequestID() string {
	return ctxProvider.ctxHTTPHead.requestID
}

func (ctxProvider ContextProvider) GetAccessKey() string {
	return ctxProvider.ctxHTTPHead.accessKey
}

func (ctxProvider ContextProvider) GetSecretKey() string {
	return ctxProvider.ctxHTTPHead.secretKey
}

func (ctxProvider ContextProvider) GetSecurityAccessKey() string {
	return ctxProvider.ctxHTTPHead.securityAccessKey
}

func (ctxProvider ContextProvider) GetSecuritySecretKey() string {
	return ctxProvider.ctxHTTPHead.securitySecretKey
}

func (ctxProvider ContextProvider) GetAuthToken() string {
	return ctxProvider.ctxHTTPHead.authToken
}

func (ctxProvider ContextProvider) GetSecurityToken() string {
	return ctxProvider.ctxHTTPHead.securityToken
}

func (ctxProvider ContextProvider) GetOriginVersionTag() string {
	return ctxProvider.ctxHTTPHead.originVersionTag
}

func (ctxProvider ContextProvider) GetWorkflowStateId() string {
	return ctxProvider.ctxHTTPHead.workflowStateId
}

func (ctxProvider ContextProvider) GetWorkflowId() string {
	return ctxProvider.ctxHTTPHead.workflowId
}

func (ctxProvider ContextProvider) GetWorkflowRunId() string {
	return ctxProvider.ctxHTTPHead.workflowRunId
}

func (ctxProvider ContextProvider) GetAlias() string {
	alias, found := strings.CutPrefix(ctxProvider.ctxHTTPHead.originVersionTag, "!")
	if found {
		return alias
	}
	return ""
}

func (ctxProvider ContextProvider) GetFuncName() string {
	// X-Cff-Func-Name returns: 0@default@go_go-doc-sample-container-event
	li := strings.LastIndex(ctxProvider.ctxHTTPHead.funcName, "@")
	if li != -1 && li+1 < len(ctxProvider.ctxHTTPHead.funcName) {
		return ctxProvider.ctxHTTPHead.funcName[li+1:]
	}
	return ctxProvider.ctxHTTPHead.funcName
}

func (ctxProvider ContextProvider) GetPackageName() string {
	return ctxProvider.ctxHTTPHead.packageName
}

func (ctxProvider ContextProvider) GetFuncVersion() string {
	return ctxProvider.ctxHTTPHead.funcVersion
}

func (ctxProvider ContextProvider) GetProjectID() string {
	return ctxProvider.ctxHTTPHead.projectID
}

func (ctxProvider ContextProvider) GetMemory() int {
	return ctxProvider.ctxHTTPHead.memory
}

func (ctxProvider ContextProvider) GetTimeout() int {
	return ctxProvider.ctxHTTPHead.timeout
}

func newContextHTTP(req *http.Request) *contextHTTP {
	if debugMode {
		log.Println("------------------ HTTP Headers ------------------")
		for name, values := range req.Header {
			log.Println(name, "=", values)
		}
		// print all environment variables
		log.Println("------------------ Environment Variables------------------")
		for _, v := range os.Environ() {
			log.Println(v)
		}
	}

	return &contextHTTP{
		region:            req.Header.Get("X-Cff-Region"),
		requestID:         req.Header.Get("X-CFF-Request-Id"),
		accessKey:         req.Header.Get("X-CFF-Access-Key"),
		secretKey:         req.Header.Get("X-CFF-Secret-Key"),
		authToken:         req.Header.Get("X-CFF-Auth-Token"),
		securityAccessKey: req.Header.Get("X-CFF-Security-Access-Key"),
		securitySecretKey: req.Header.Get("X-CFF-Security-Secret-Key"),
		securityToken:     req.Header.Get("X-CFF-Security-Token"),
		originVersionTag:  req.Header.Get("X-Cff-Origin-Version-Tag"),
		workflowStateId:   req.Header.Get("X-Cff-Workflow-State-Id"),
		workflowId:        req.Header.Get("X-Cff-Workflow-Id"),
		workflowRunId:     req.Header.Get("X-Cff-Workflow-Run-Id"),
		funcName:          req.Header.Get("X-Cff-Func-Name"),
		packageName:       req.Header.Get("X-Cff-Package"),
		funcVersion:       req.Header.Get("X-Cff-Func-Version"),
		projectID:         req.Header.Get("X-Cff-Project-Id"),
		timeout:           getIntHeader(req.Header, "X-Cff-Timeout"),
		memory:            getIntHeader(req.Header, "X-Cff-Memory"),
		fcStartTime:       getCurrentTime(),
	}
}
