package context

import (
	"github.com/opentelekomcloud-community/otc-functiongraph-go-runtime/go-runtime/pkg/runtime/common"
)

// The Context interface allows functions to obtain the function execution context,
// such as the AccessKey/SecretKey delegated by the user, the current request ID,
// the memory space allocated for function execution, the number of CPUs, and so on.

type RuntimeContext interface {
	// Get the request ID.
	GetRequestID() string

	// Get the remaining running time of the function.
	GetRemainingTimeInMilliSeconds() int

	// Get the AccessKey delegated by the user (valid for 24 hours).
	// To use this method, you need to configure delegation for the function.
	// The current function workflow has stopped maintaining the getAccessKey
	// interface in the Runtime SDK.
	// You will not be able to use getAccessKey to obtain a temporary AK.
	GetAccessKey() string

	// Get the SecretKey of the user's delegation (valid for 24 hours).
	// To use this method, you need to configure the delegation for the function.
	// The current function workflow has stopped maintaining the getSecretKey
	// interface in the Runtime SDK.
	// You will not be able to use getSecretKey to obtain a temporary SK.
	GetSecretKey() string

	// Get the SecurityAccessKey delegated by the user (valid for 24 hours).
	// To use this method, you need to configure a delegate for the function.
	GetSecurityAccessKey() string

	// Get the SecuritySecretKey (valid for 24 hours) delegated by the user.
	// To use this method, you need to configure the delegate for the function.
	GetSecuritySecretKey() string

	// Get the function name.
	GetFunctionName() string

	// Get the value passed by the user through the environment variable through key.
	GetUserData(string) string

	GetLogger() common.RuntimeLogger

	// Get the function timeout.
	GetRunningTimeInSeconds() int

	// Get the version of the function.
	GetVersion() string

	//Get allocated memory.
	GetMemorySize() int

	// Get the CPU resources used by the function.
	GetCPUNumber() int

	// Get project id
	GetProjectID() string

	// Gets a function group.
	GetPackage() string

	// Get the user's delegated token (valid for 24 hours).
	// To use this method, you need to configure the delegate for the function.
	GetToken() string

	// Get the SecurityToken delegated by the user (valid for 24 hours).
	// To use this method, you need to configure the delegate for the function.
	GetSecurityToken() string

	// Get the alias of a function
	GetAlias() string

	GetStreamRunID() string

	GetStreamEnable() string

	GetStreamToken() string

	GetStreamAddr() string

	GetOriginVersionTag() string

	GetWorkflowStateID() string

	GetWorkflowID() string

	GetWorkflowRunID() string
	

	
	// Get the alias from Header
	GetAlias2() string
}
