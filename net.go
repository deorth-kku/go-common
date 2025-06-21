package common

import (
	"net"
	"net/netip"
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

func AddrFromSlice(ip []byte) netip.Addr {
	ip = ShortIP(ip)
	switch len(ip) {
	case net.IPv4len:
		return netip.AddrFrom4([4]byte(ip))
	case net.IPv6len:
		return netip.AddrFrom16([16]byte(ip))
	default:
		return netip.Addr{}
	}
}

func ParseIP(ip string) net.IP {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil
	}
	return addr.AsSlice()
}
