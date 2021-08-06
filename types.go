package common

import (
	"github.com/anacrolix/dht/v2/krpc"
	"net"
)

// A collection of MonitoredNodeJSON
type MonitorOutputJSON struct {
	StartDate string               `json:"start_timestamp"`
	EndDate   string               `json:"end_timestamp"`
	Nodes     []*MonitoredNodeJSON `json:"found_nodes"`
}

// A DHT Node described by its UDP address, its connectivity status and IP Version encoded as JSON strings
type MonitoredNodeJSON struct {
	UDPAddrs  net.UDPAddr `json:"udpaddrs"`
	Reachable bool        `json:"reachable"`
	IPVersion string      `json:"ip_version"`
}

// Monitor Output object that contains the results of the
type Output struct {
	StartDate string
	EndDate   string
	Nodes     map[string]*ObservedNode
}

// All the information about a single node during monitoring
type ObservedNode struct {
	UDPAddr   net.UDPAddr
	Reachable bool
	IPVersion IPVersion
}

// IPVersion is the version of an IP address, or unknown for invalid addresses.
type IPVersion uint

const (
	UnknownVersion IPVersion = iota
	IPv4
	IPv6
)

// KRPCResponseError signals that the KRPC Query returned an error. These could be:
// 1. A nil pointer in the message
// 2. Maybe a anacrolix/dht's rate limit error (TBD)
type KRPCNilResponseError struct {
	Msg  string
	Peer krpc.NodeInfo
}
