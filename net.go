package common

import (
	"net"
	"strconv"
)

func JoinHostPort[T AnyInt](host string, port T) string {
	return net.JoinHostPort(host, strconv.FormatUint(uint64(port), 10))
}

func ShortIP(ip net.IP) net.IP {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4
	}
	return ip
}

func ParseIP(ip string) net.IP {
	return ShortIP(net.ParseIP(ip))
}
