package mainline_dht_common

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"

	dht "github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/krpc"
)

func TestParseAddr(t *testing.T) {
	bootstrap := "router.bittorrent.com:6881"
	addr := &krpc.NodeAddr{
		IP:   net.ParseIP("67.215.246.10"),
		Port: 6881,
	}
	actual, _ := ParseAddrString(bootstrap)
	expected := &krpc.NodeInfo{
		ID:   dht.RandomNodeID(),
		Addr: *addr,
	}
	assert.Equal(t, actual.Addr.Port, expected.Addr.Port, "expected: expected: %d -- actual: %d", expected.Addr.Port, actual.Addr.Port)
	assert.True(t, expected.Addr.IP.Equal(actual.Addr.IP), "expected: expected: %s -- actual: %s ", expected.Addr.IP.String(), actual.Addr.IP.String())
}
