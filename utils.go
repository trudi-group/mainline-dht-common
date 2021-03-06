package common

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"

	dht "github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/int160"
	"github.com/anacrolix/dht/v2/krpc"
)

// IsNullSlice checks if the bytes of given slice are all null
func IsNullSlice(arr [20]byte) bool {
	emptySlice := make([]byte, 20)
	val := make([]byte, 20)
	copy(val, arr[:])
	if bytes.Equal(val, emptySlice) {
		return true
	}
	return false
}

// AskYesNo asks the user to enter "y" or "n". Also recognizes "yes" and "no" in all capitalizations.
// Anything unrecognized will is equivalent to "n".
func AskYesNo() bool {
	var response string
	positiveResp := []string{"y", "yes"}
	// negativeResp := []string{"n", "no"}

	_, err := fmt.Scanln(&response)
	if err != nil {
		return false
	}

	response = strings.ToLower(response)
	if containsOneResponse(response, positiveResp) {
		return true
	}
	return false
}

func containsOneResponse(inputString string, resp []string) bool {
	for _, r := range resp {
		if strings.Contains(inputString, r) {
			return true
		}
	}
	return false
}

// ParseAddrString parses an IP:port address into anacrolix/dht krpc.NodeInfo. This is very useful when connecting to bootstrap nodes.
func ParseAddrString(text string) (*krpc.NodeInfo, error) {
	host, port, err := net.SplitHostPort(text)
	var portAsInt int
	var nodeInfo *krpc.NodeInfo
	if err != nil {
		return nil, err
	}
	portAsInt, err = strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	id := dht.RandomNodeID()
	addresses, err := net.LookupIP(host)
	if err == nil {
		ip := addresses[0]
		addr := &krpc.NodeAddr{
			IP:   ip,
			Port: portAsInt,
		}

		nodeInfo = &krpc.NodeInfo{
			ID:   id,
			Addr: *addr,
		}
		return nodeInfo, nil
	}
	return nil, err
}

// TargetIsInZone returns true is targetID has zone many bits in common with sourceID
func TargetIsInZone(sourceID [20]byte, targetID [20]byte, zone int) bool {
	source := int160.FromByteArray(sourceID)
	target := int160.FromByteArray(targetID)
	var dist int160.T
	dist.Xor(&source, &target)
	for i := 0; i < 160; i++ {
		if i >= zone {
			// Got passed the zone with only 0 bits
			return true
		} else {
			// Found a set bit among the first zone many bits
			if dist.GetBit(i) {
				return false
			}
		}
	}
	return false
}

// NodeInfoToID extracts each ID from an arry of krpc.NodeInfo and returns the krpc.IDs
func NodeInfoToID(addrs []*krpc.NodeInfo) []krpc.ID {
	peers := make([]krpc.ID, len(addrs))
	for i, addr := range addrs {
		peers[i] = addr.ID
	}
	return peers
}

// UDPToNodeAddr converts the net.UDPAddr to a krpc.NodeAddr
func UDPToNodeAddr(udp net.UDPAddr) krpc.NodeAddr {
	var addr krpc.NodeAddr
	addr.FromUDPAddr(&udp)
	return addr
}

// NodeinfoToUdpAddr converts a krpc.NodeInfo to a net.UDPAddr
func NodeinfoToUdpAddr(nodeInfo *krpc.NodeInfo) net.UDPAddr {
	var udp net.UDPAddr
	udp = *nodeInfo.Addr.UDP()
	return udp
}
