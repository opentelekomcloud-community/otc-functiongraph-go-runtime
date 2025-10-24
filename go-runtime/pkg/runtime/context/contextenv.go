package rtcontext

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	defaultTimeout = "3"
)

func atoi(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		log.Println("execute strconv.Atoi failed.")
	}
	return result
}

func getCurrentTime() int64 {
	return (time.Now().UnixNano() / 1000000)
}

func (contextObj *ContextEnv) InitializeContext() {
	timeout := os.Getenv("RUNTIME_TIMEOUT")
	if timeout == "" {
		timeout = defaultTimeout
	}
	contextObj.rtTimeout = atoi(timeout)
	rtProjectID := os.Getenv("RUNTIME_PROJECT_ID")
	if rtProjectID != "" {
		contextObj.rtProjectID = rtProjectID
	}
	rtPackage := os.Getenv("RUNTIME_PACKAGE")
	if rtPackage != "" {
		contextObj.rtPackage = rtPackage
	}
	rtFcName := os.Getenv("RUNTIME_FUNC_NAME")
	if rtFcName != "" {
		contextObj.rtFcName = rtFcName
	}

	rtAlias := os.Getenv("RUNTIME_ALIAS")
	if rtAlias != "" {
		contextObj.rtAlias = rtAlias
	}

	rtFcVersion := os.Getenv("RUNTIME_FUNC_VERSION")
	if rtFcVersion != "" {
		contextObj.rtFcVersion = rtFcVersion
	}
	rtMemory := os.Getenv("RUNTIME_MEMORY")
	if rtMemory != "" {
		contextObj.rtMemory = atoi(rtMemory)
	}
	rtCPU := os.Getenv("RUNTIME_CPU")
	if rtCPU != "" {
		contextObj.rtCPU = atoi(rtCPU)
	}
	rtUserData := os.Getenv("RUNTIME_USERDATA")
	if rtUserData != "" {
		err := json.Unmarshal([]byte(rtUserData), &contextObj.rtUserData)
		if err != nil {
			log.Println("Failed to Unmarshal Userdata")
		}
	}
	contextObj.rtHandler = os.Getenv("RUNTIME_HANDLER")
	InitializerHandler := os.Getenv("RUNTIME_INITIALIZER_HANDLER")
	if InitializerHandler != "" {
		contextObj.rtInitializerHandler = InitializerHandler
		InitializerTimeout := os.Getenv("RUNTIME_INITIALIZER_TIMEOUT")
		if InitializerTimeout == "" {
			InitializerTimeout = defaultTimeout
		}
		contextObj.rtInitializerTimeout = atoi(InitializerTimeout)
	}

	rtMaxRequestBodySize := os.Getenv("RUNTIME_MAX_REQUEST_BODY_SIZE")
	if rtMaxRequestBodySize != "" {
		contextObj.rtMaxRequestBodySize = atoi(rtMaxRequestBodySize)
	}
	rtPreStopHandler := os.Getenv("RUNTIME_PRESTOP_HANDLER")
	if rtPreStopHandler != "" {
		contextObj.rtPreStopHandler = rtPreStopHandler
	}
	rtRunMaxStateSize := os.Getenv("RUNTIME_MAX_STATE_SIZE")
	if rtRunMaxStateSize != "" {
		contextObj.rtRunMaxStateSize = atoi(rtRunMaxStateSize)
	}
	rtMaxResponseBodySize := os.Getenv("RUNTIME_MAX_RESP_BODY_SIZE")
	if rtMaxResponseBodySize != "" {
		contextObj.rtMaxResponseBodySize = atoi(rtMaxResponseBodySize)
	}
}
