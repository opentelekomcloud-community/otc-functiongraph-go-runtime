package rtcontext

type ContextEnv struct {
	rtProjectID          string            // RUNTIME_PROJECT_ID
	rtFcName             string            // RUNTIME_FUNC_NAME
	rtFcVersion          string            // RUNTIME_FUNC_VERSION
	rtPackage            string            // RUNTIME_PACKAGE
	rtMemory             int               // RUNTIME_MEMORY
	rtCPU                int               // RUNTIME_CPU
	rtTimeout            int               // RUNTIME_TIMEOUT
	rtHandler            string            // RUNTIME_HANDLER
	rtUserData           map[string]string // RUNTIME_USER_DATA
	rtInitializerTimeout int               // RUNTIME_INITIALIZER_TIMEOUT
	rtInitializerHandler string            // RUNTIME_INITIALIZER_HANDLER
	rtAlias              string            // RUNTIME_ALIAS
	// new
	rtMaxRequestBodySize  int    // RUNTIME_MAX_REQUEST_BODY_SIZE
	rtPreStopHandler      string // RUNTIME_PRESTOP_HANDLER
	rtRunMaxStateSize     int    // RUNTIME_MAX_STATE_SIZE
	rtMaxResponseBodySize int    // RUNTIME_MAX_RESP_BODY_SIZE
}

type ContextHTTP struct {
	requestID         string
	accesskey         string
	secretKey         string
	securityAccessKey string
	securitySecretKey string
	authToken         string
	fcStartTime       int64
	rtRemainTime      int
	securityToken     string
	// new
	streamRunId      string
	streamEnable     string
	streamToken      string
	streamAddr       string
	originVersionTag string
	workflowStateId  string
	workflowId       string
	workflowRunId    string
}

type ContextProvider struct {
	ctxEnv      *ContextEnv
	ctxHTTPHead *ContextHTTP
}

type userFunctionLog struct {
	requestID string
}

func (l *userFunctionLog) getRequestId() string {
	return l.requestID
}

func (l *userFunctionLog) setRequestId(requestId string) {
	l.requestID = requestId
}
