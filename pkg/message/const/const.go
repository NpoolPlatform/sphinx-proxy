package constant

import (
	"strings"
	"time"
)

const (
	GrpcTimeout = time.Second * 10
)

const (
	ServiceName = "sphinx-proxy.npool.top"
)

func FormatServiceName() string {
	return strings.Title(ServiceName)
}
