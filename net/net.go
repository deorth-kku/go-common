package cnet

import (
	"net"
	"net/netip"
	"strconv"

	"github.com/deorth-kku/go-common/args"
	cmath "github.com/deorth-kku/go-common/math"
)

func JoinHostPort[T cmath.AnyInt](host string, port T) string {
	return net.JoinHostPort(host, strconv.FormatUint(uint64(port), 10))
}

func ShortIP(ip net.IP) net.IP {
	if ip4 := ip.To4(); ip4 != nil {
		return ip4
	}
	return ip
}

func AddrFromSlice(ip []byte) netip.Addr {
	return args.Drop1(netip.AddrFromSlice(ip))
}

func ParseIP(ip string) net.IP {
	addr, err := netip.ParseAddr(ip)
	if err != nil {
		return nil
	}
	return addr.AsSlice()
}
