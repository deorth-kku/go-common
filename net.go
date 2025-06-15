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
	if ip4 := net.IP(ip).To4(); ip4 != nil {
		return netip.AddrFrom4([4]byte(ip4))
	}
	return netip.AddrFrom16([16]byte(net.IP(ip).To16()))
}

func ParseIP(ip string) net.IP {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil
	}
	return addr.AsSlice()
}
