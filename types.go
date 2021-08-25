package common

import (
	"github.com/anacrolix/dht/v2/int160"
	"github.com/anacrolix/dht/v2/krpc"
	"net"
)

// MonitoredNodeJSON is a collection of MonitoredNodeJSON
type MonitorOutputJSON struct {
	StartDate string               `json:"start_timestamp"`
	EndDate   string               `json:"end_timestamp"`
	Nodes     []*MonitoredNodeJSON `json:"found_nodes"`
}

// MonitoredNodeJSON holds a DHT Node described by its UDP address, its connectivity status and IP Version encoded as JSON strings
type MonitoredNodeJSON struct {
	UDPAddrs  net.UDPAddr `json:"udpaddrs"`
	Reachable bool        `json:"reachable"`
	IPVersion string      `json:"ip_version"`
}

// Output is an object that contains the results of the monitoring period, i.e. start and end time, and a map of ObservedNodes
type Output struct {
	StartDate string
	EndDate   string
	Nodes     map[string]*ObservedNode
}

// ObservedNode contains all the information about a single node during observation
type ObservedNode struct {
	UDPAddr   net.UDPAddr
	Reachable bool
	IPVersion IPVersion
}

// IPVersion is the version of an IP address, or unknown for invalid addresses.
type IPVersion uint

// Enum describing possible types of IP addresses
const (
	UnknownVersion IPVersion = iota
	IPv4
	IPv6
)

// KRPCNilResponseError signals that the KRPC Query returned an error. These could be:
// 1. A nil pointer in the message
// 2. Maybe a anacrolix/dht's rate limit error (TBD)
type KRPCNilResponseError struct {
	Msg  string
	Peer krpc.NodeInfo
}

func (e *KRPCNilResponseError) Error() string {
	return e.Msg
}

// CrawlOutput summarises the crawl, i.e. timestamp + collection of nodes we learned
type CrawlOutput struct {
	StartDate string
	EndDate   string
	Nodes     map[krpc.ID]*CrawledNode
}

// CrawledNode stores everything we know about a node we have contacted
type CrawledNode struct {
	NID        krpc.ID
	UDPAddrs   net.UDPAddr
	Reachable  bool
	IPVersion  IPVersion
	Neighbours []krpc.ID
	Timestamp  string
}

// CrawlResult is a container struct for crawl results... because of go...
type CrawlResult struct {
	Node *NodeKnows
	Err  error
}

// NodeKnows stores the collected addresses for a given ID
type NodeKnows struct {
	id    krpc.ID
	knows []*krpc.NodeInfo
	info  map[string]interface{}
}

// InfoHash is a 20 byte ID
type InfoHash = int160.T
