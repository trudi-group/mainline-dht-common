package common

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetermineIPVersion(t *testing.T) {
	tests := []struct {
		name     string
		ip       net.IP
		expected IPVersion
	}{
		{
			// The standard library parses IPv4 addresses to 16-byte
			// "IPv4-mapped IPv6 addresses", see RFC 4291 Sec. 2.5.5.2.
			name:     "parsed-IPv4",
			ip:       net.ParseIP("127.0.0.1"),
			expected: IPv4,
		},
		{
			name:     "4-byte-IPv4",
			ip:       net.IP{127, 0, 0, 1},
			expected: IPv4,
		},
		{
			// This also produces an IPv4-mapped IPv6 address.
			name:     "net-IPv4",
			ip:       net.IPv4(127, 0, 0, 1),
			expected: IPv4,
		},
		{
			name:     "Invalid",
			ip:       net.IP{127},
			expected: UnknownVersion,
		},
		{
			name:     "IPv6",
			ip:       net.ParseIP("::FFFF:C0A8"),
			expected: IPv6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := DetermineIPVersion(test.ip)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsIPAddressLocal(t *testing.T) {
	tests := []struct {
		name     string
		ip       string
		expected bool
	}{
		{
			name:     "IPv4-loopback",
			ip:       "127.0.0.1",
			expected: true,
		},
		{
			name:     "IPv4-private-1",
			ip:       "192.168.2.3",
			expected: true,
		},
		{
			name:     "IPv4-private-2",
			ip:       "10.12.14.15",
			expected: true,
		},
		{
			name:     "IPv4-private-3",
			ip:       "172.16.33.33",
			expected: true,
		},
		{
			name:     "IPv4-carrier-grade-NAT",
			ip:       "100.64.12.13",
			expected: true,
		},
		{
			name:     "IPv4-public",
			ip:       "1.2.3.4",
			expected: false,
		},
		{
			name:     "IPv6-loopback",
			ip:       "::1",
			expected: true,
		},
		{
			name:     "IPv6-ULA",
			ip:       "fd00::1",
			expected: true,
		},
		{
			name:     "IPv6-public",
			ip:       "c0ff:eeee::1",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsIPAddressLocal(net.ParseIP(test.ip))
			assert.Equal(t, test.expected, actual)
		})
	}
}
