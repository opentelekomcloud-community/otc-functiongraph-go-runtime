package rtcontext

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime/common"
)

var (
	once       sync.Once
	contextobj *ContextEnv

	funcLogger = &userFunctionLog{}
)

var DEBUG_MODE = false

func (ctxProvider ContextProvider) GetRemainingTimeInMilliSeconds() int {
	currentTime := getCurrentTime()
	usedTime := int(currentTime - ctxProvider.ctxHTTPHead.fcStartTime)
	timeout := ctxProvider.ctxEnv.rtTimeout * 1000
	if usedTime < timeout {
		ctxProvider.ctxHTTPHead.rtRemainTime = timeout - usedTime
	} else {
		ctxProvider.ctxHTTPHead.rtRemainTime = 0
	}
	return ctxProvider.ctxHTTPHead.rtRemainTime
}

func (ctxProvider ContextProvider) GetFunctionName() string {
	return ctxProvider.ctxEnv.rtFcName
}

func (ctxProvider ContextProvider) GetRunningTimeInSeconds() int {
	return ctxProvider.ctxEnv.rtTimeout
}

func (ctxProvider ContextProvider) GetVersion() string {
	return ctxProvider.ctxEnv.rtFcVersion
}

func (ctxProvider ContextProvider) GetMemorySize() int {
	return ctxProvider.ctxEnv.rtMemory
}

func (ctxProvider ContextProvider) GetCPUNumber() int {
	return ctxProvider.ctxEnv.rtCPU
}

func (ctxProvider ContextProvider) GetUserData(key string) string {
	if ctxProvider.ctxEnv.rtUserData != nil {
		return ctxProvider.ctxEnv.rtUserData[key]
	}
	return ""
}

func (ctxProvider ContextProvider) GetLogger() common.RuntimeLogger {
	if funcLogger.getRequestId() != ctxProvider.ctxHTTPHead.requestID {
		funcLogger.setRequestId(ctxProvider.ctxHTTPHead.requestID)
	}

	return funcLogger
}

func (ctxProvider ContextProvider) GetProjectID() string {
	return ctxProvider.ctxEnv.rtProjectID
}

func (ctxProvider ContextProvider) GetPackage() string {
	return ctxProvider.ctxEnv.rtPackage
}

func (ctxProvider ContextProvider) GetHandler() string {
	return ctxProvider.ctxEnv.rtHandler
}

func (ctxProvider ContextProvider) GetInitializerHandler() string {
	return ctxProvider.ctxEnv.rtInitializerHandler
}

func (ctxProvider ContextProvider) GetAlias() string {
	return ctxProvider.ctxEnv.rtAlias
}

func (ctxProvider ContextProvider) GetMaxRequestBodySize() int {
	return ctxProvider.ctxEnv.rtMaxRequestBodySize
}

func (ctxProvider ContextProvider) GetPreStopHandler() string {
	return ctxProvider.ctxEnv.rtPreStopHandler
}

func (ctxProvider ContextProvider) GetRunMaxStateSize() int {
	return ctxProvider.ctxEnv.rtRunMaxStateSize
}

func (ctxProvider ContextProvider) GetMaxResponseBodySize() int {
	return ctxProvider.ctxEnv.rtMaxResponseBodySize
}

func (ctxProvider ContextProvider) GetAccessKey() string {
	return ctxProvider.ctxHTTPHead.accesskey
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

func (ctxProvider ContextProvider) GetToken() string {
	return ctxProvider.ctxHTTPHead.authToken
}

func (ctxProvider ContextProvider) GetRequestID() string {
	return ctxProvider.ctxHTTPHead.requestID
}

func (ctxProvider ContextProvider) GetSecurityToken() string {
	return ctxProvider.ctxHTTPHead.securityToken
}

func (ctxProvider ContextProvider) GetStreamRunID() string {
	return ctxProvider.ctxHTTPHead.streamRunId
}

func (ctxProvider ContextProvider) GetStreamEnable() string {
	return ctxProvider.ctxHTTPHead.streamEnable
}

func (ctxProvider ContextProvider) GetStreamToken() string {
	return ctxProvider.ctxHTTPHead.streamToken
}

func (ctxProvider ContextProvider) GetStreamAddr() string {
	return ctxProvider.ctxHTTPHead.streamAddr
}

func (ctxProvider ContextProvider) GetOriginVersionTag() string {
	return ctxProvider.ctxHTTPHead.originVersionTag
}

func (ctxProvider ContextProvider) GetAlias2() string {
	alias, found := strings.CutPrefix(ctxProvider.ctxHTTPHead.originVersionTag, "!")
	if found {
		return alias
	}
	return ""
}

func (ctxProvider ContextProvider) GetWorkflowStateID() string {
	return ctxProvider.ctxHTTPHead.workflowStateId
}

func (ctxProvider ContextProvider) GetWorkflowID() string {
	return ctxProvider.ctxHTTPHead.workflowId
}

func (ctxProvider ContextProvider) GetWorkflowRunID() string {
	return ctxProvider.ctxHTTPHead.workflowRunId
}

func (logger *userFunctionLog) Logf(format string, args ...interface{}) {
	// logTimeFormat := "2006-01-02 15:04:05.999-07:00"
	logTimeFormat := "2006-01-02T15:04:05.000Z"
	content := fmt.Sprintf(format, args...)
	myFormat := fmt.Sprintf("%s %s %s", time.Now().UTC().Format(logTimeFormat), logger.requestID, content)
	fmt.Println(myFormat)
}

func GetContextProvider(ctxEnv *ContextEnv, ctxHTTPHead *ContextHTTP) ContextProvider {
	return ContextProvider{ctxEnv, ctxHTTPHead}
}

func GetContextEnvInstance() *ContextEnv {

	once.Do(func() {
		contextobj = new(ContextEnv)
	})
	return contextobj
}

func GetContextHTTPHeadInstance(req *common.InvokeRequest) *ContextHTTP {
	contextHTTPHead := new(ContextHTTP)

	if DEBUG_MODE {
		for name, values := range req.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(name, value)
			}
		}
	}

	requestID := req.Header.Get("X-CFF-Request-Id")
	if requestID != "" {
		contextHTTPHead.requestID = requestID
	}
	accesskey := req.Header.Get("X-CFF-Access-Key")
	if accesskey != "" {
		contextHTTPHead.accesskey = accesskey
	}
	secretKey := req.Header.Get("X-CFF-Secret-Key")
	if secretKey != "" {
		contextHTTPHead.secretKey = secretKey
	}
	authToken := req.Header.Get("X-CFF-Auth-Token")
	if authToken != "" {
		contextHTTPHead.authToken = authToken
	}
	contextHTTPHead.fcStartTime = getCurrentTime()
	securityAccessKey := req.Header.Get("X-CFF-Security-Access-Key")
	if securityAccessKey != "" {
		contextHTTPHead.securityAccessKey = securityAccessKey
	}
	securitySecretKey := req.Header.Get("X-CFF-Security-Secret-Key")
	if securitySecretKey != "" {
		contextHTTPHead.securitySecretKey = securitySecretKey
	}
	securityToken := req.Header.Get("X-CFF-Security-Token")
	if securityToken != "" {
		contextHTTPHead.securityToken = securityToken
	}

	streamRunId := req.Header.Get("X-Stream-Run-Id")
	if streamRunId != "" {
		contextHTTPHead.streamRunId = streamRunId
	}
	streamEnable := req.Header.Get("X-Stream-Enable")
	if streamEnable != "" {
		contextHTTPHead.streamEnable = streamEnable
	}
	streamToken := req.Header.Get("X-Stream-Token")
	if streamToken != "" {
		contextHTTPHead.streamToken = streamToken
	}
	streamAddr := req.Header.Get("X-Stream-Addr")
	if streamAddr != "" {
		contextHTTPHead.streamAddr = streamAddr
	}

	originVersionTag := req.Header.Get("X-Cff-Origin-Version-Tag")
	if originVersionTag != "" {
		contextHTTPHead.originVersionTag = originVersionTag
	}

	workflowStateId := req.Header.Get("X-Cff-Workflow-State-Id")
	if workflowStateId != "" {
		contextHTTPHead.workflowStateId = workflowStateId
	}

	workflowId := req.Header.Get("X-Cff-Workflow-Id")
	if workflowId != "" {
		contextHTTPHead.workflowId = workflowId
	}

	workflowRunId := req.Header.Get("X-Cff-Workflow-Run-Id")
	if workflowRunId != "" {
		contextHTTPHead.workflowRunId = workflowRunId
	}

	return contextHTTPHead
}
