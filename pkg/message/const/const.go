package constant

import (
	"strings"
	"time"
)

const (
	// GrpcTimeout ..
	GrpcTimeout = time.Second * 10
	// ServiceName ..
	ServiceName = "sphinx-proxy.npool.top"

	// DefaultPageSize ..
	DefaultPageSize = 100
	// TaskDuration ..
	TaskDuration = time.Second * 5
	// TaskTimeout ..
	TaskTimeout = time.Second * 10
)

// FormatServiceName ..
func FormatServiceName() string {
	return strings.Title(ServiceName)
}
