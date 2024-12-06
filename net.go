package common

import (
	"fmt"
	"net"
)

func JoinHostPort[T AnyInt](host string, port T) string {
	return net.JoinHostPort(host, fmt.Sprintf("%d", port))
}
