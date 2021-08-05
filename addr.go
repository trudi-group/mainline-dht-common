// Provided by @mrd0ll4r in https://github.com/cndolo/mainline-dht-crawler

package mainline_dht_common

import (
	"fmt"
	"net"
)

func (v IPVersion) String() string {
	switch v {
	case IPv4:
		return "IPv4"
	case IPv6:
		return "IPv6"
	case UnknownVersion:
		return "Unknown"
	default:
		panic(fmt.Sprintf("invalid IPVersion: %d", v))
	}
}

// DetermineIPVersion determines whether the given address is v4, v6, or
// invalid.
// This essentially duplicates the logic in (net.IP).To4() and (net.IP).To16()
// without the special-casing of IPv4-in-IPv6 addresses.
func DetermineIPVersion(address net.IP) IPVersion {
	if address.To4() != nil {
		return IPv4
	} else if address.To16() != nil {
		return IPv6
	} else {
		return UnknownVersion
	}
}

var ipv4localIPNets = []net.IPNet{
	// IPv4 loopback
	{
		IP:   net.IPv4(127, 0, 0, 0),
		Mask: net.CIDRMask(8, 8*net.IPv4len),
	},
	// IPv4 local stuff
	{
		IP:   net.IPv4(0, 0, 0, 0),
		Mask: net.CIDRMask(8, 8*net.IPv4len),
	},
	// IPv4 private use
	{
		IP:   net.IPv4(10, 0, 0, 0),
		Mask: net.CIDRMask(8, 8*net.IPv4len),
	},
	// IPv4 private use
	{
		IP:   net.IPv4(172, 16, 0, 0),
		Mask: net.CIDRMask(12, 8*net.IPv4len),
	},
	// IPv4 private use
	{
		IP:   net.IPv4(192, 168, 0, 0),
		Mask: net.CIDRMask(16, 8*net.IPv4len),
	},
	// IPv4 carrier-grade NAT stuff
	{
		IP:   net.IPv4(100, 64, 0, 0),
		Mask: net.CIDRMask(10, 8*net.IPv4len),
	},
}
var ipv6LocalIPNets = []net.IPNet{
	// IPv6 loopback
	{
		IP:   net.IPv6loopback,
		Mask: net.CIDRMask(128, 8*net.IPv6len),
	},
	// IPv6 Unique Local Addresses
	{
		IP:   net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(7, 8*net.IPv6len),
	},
}

// IsIPAddressLocal determines whether an IP address is local.
// Local addresses should(?) not be routed on the open internet and are probably
// not connectable.
func IsIPAddressLocal(addr net.IP) bool {
	if addr == nil {
		panic("nil IP")
	}
	if a := addr.To4(); a != nil {
		for _, subnet := range ipv4localIPNets {
			if subnet.Contains(a) {
				return true
			}
		}
	} else {
		for _, subnet := range ipv6LocalIPNets {
			if subnet.Contains(addr) {
				return true
			}
		}
	}
	return false
}
