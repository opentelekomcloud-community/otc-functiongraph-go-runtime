package invoke_fg

import (
	"testing"
)

// to run the test:
// go test -run TestInvokeSync_AKSK
func TestInvokeSync_AKSK(t *testing.T) {
	InvokeSync_AKSK()
}

// to run the test:
// go test -run TestInvokeSync_UsernamePassword
func TestInvokeSync_UsernamePassword(t *testing.T) {
	InvokeSync_UsernamePassword()
}

// to run the test:
// go test -run TestInvokeAsync_UsernamePassword
func TestInvokeAsync_UsernamePassword(t *testing.T) {
	InvokeAsync_UsernamePassword()
}
