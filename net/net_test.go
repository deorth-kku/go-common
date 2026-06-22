package cnet

import (
	"net"
	"testing"
)

var ips = []string{
	"192.168.1.1",
	"10.0.0.0",
	"172.16.254.1",
	"2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	"::1",
	"fe80::1ff:fe23:4567:890a",
	"2001:db8::",
	"::ffff:192.0.2.128",
}

func getip(count int) string {
	return ips[count%len(ips)]
}

func BenchmarkOldParseIP(b *testing.B) {
	for i := range b.N {
		ShortIP(net.ParseIP(getip(i)))
	}
}

func BenchmarkParseIP(b *testing.B) {
	for i := range b.N {
		ParseIP(getip(i))
	}
}

func TestParseIP(t *testing.T) {
	for i, ip := range ips {
		parsed := ParseIP(ip)
		if parsed == nil {
			t.Errorf("failed to parse IP %s at index %d", ip, i)
		}
	}
}
