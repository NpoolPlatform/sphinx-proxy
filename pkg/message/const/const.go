package constant

import (
	"time"
)

const (
	// GrpcTimeout ..
	GrpcTimeout = time.Second * 10

	// ServiceName ..
	ServiceName = "sphinx-proxy.npool.top"

	// DefaultPageSize ..
	DefaultPageSize = 50
	// TaskDuration ..
	TaskDuration = time.Second * 5
	// TaskTimeout ..
	TaskTimeout = time.Second * 10
)
