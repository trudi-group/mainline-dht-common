package mainline_dht_common

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"

	dht "github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/krpc"
)

func IsNullSlice(arr [20]byte) bool {
	emptySlice := make([]byte, 20)
	val := make([]byte, 20)
	copy(val, arr[:])
	if bytes.Equal(val, emptySlice) {
		return true
	} else {
		return false
	}
}

// Asks the user to enter "y" or "n". Also recognizes "yes" and "no" in all capitalizations.
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
	} else {
		return false
	}
}

func containsOneResponse(inputString string, resp []string) bool {
	for _, r := range resp {
		if strings.Contains(inputString, r) {
			return true
		}
	}
	return false
}

// Parses an IP:port address into anacrolix/dht krpc.NodeInfo. This is very useful when connecting to bootstrap nodes.
func ParseAddrString(text string) (*krpc.NodeInfo, error) {
	host, port, err := net.SplitHostPort(text)
	var portAsInt int
	var nodeInfo *krpc.NodeInfo
	nodeInfo = nil
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
	} else {
		return nil, err
	}
}

func (e *KRPCNilResponseError) Error() string {
	return e.Msg
}
