package constant

import (
	"strings"
	"time"
)

const (
	GrpcTimeout = time.Second * 10 // nolint
)

const (
	ServiceName = "sphinx-proxy.npool.top" // nolint
)

func FormatServiceName() string {
	return strings.Title(ServiceName)
}
