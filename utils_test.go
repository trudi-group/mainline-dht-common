package common

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"

	dht "github.com/anacrolix/dht/v2"
	"github.com/anacrolix/dht/v2/int160"
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

func TestInfohashDistance(t *testing.T) {
	const zeroID = "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	zeroInfohash := int160.FromByteString(zeroID)
	for i := 12; i < 15; i++ {
		zeroInfohash.SetBit(i, true)
	}
	boundary := zeroInfohash
	zeroInfohash.SetBit(11, true)
	justUnder := zeroInfohash
	randomID := int160.FromByteArray(dht.RandomNodeID())
	samePrefixRandomID := int160.FromByteArray(dht.RandomNodeID())
	for i := 0; i < 12; i++ {
		samePrefixRandomID.SetBit(i, randomID.GetBit(i))
	}
	tests := []struct {
		name     string
		source   [20]byte
		target   [20]byte
		zone     int
		expected bool
	}{
		{
			name:     "Same-Zero-ID",
			source:   krpc.IdFromString("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			target:   krpc.IdFromString("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			expected: true,
		},
		{
			name:     "Max-distance",
			source:   krpc.IdFromString("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			target:   krpc.IdFromString("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			zone:     12,
			expected: false,
		},
		{
			name:     "1-Bit-Distance",
			source:   krpc.IdFromString("00000000000000000000"),
			target:   krpc.IdFromString("00000000000000000001"),
			zone:     12,
			expected: true,
		},
		{
			name:     "Out-of-zone",
			source:   krpc.IdFromString("11111111110111111111"),
			target:   krpc.IdFromString("\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11"),
			zone:     12,
			expected: false,
		},
		{
			name:     "Zone-boundary",
			source:   krpc.IdFromString("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			target:   boundary.AsByteArray(),
			zone:     12,
			expected: true,
		},
		{
			name:     "in 11-bit-zone",
			source:   krpc.IdFromString("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			target:   justUnder.AsByteArray(),
			zone:     12,
			expected: false,
		},
		{
			name:     "Random-ID-With-CP",
			source:   randomID.AsByteArray(),
			target:   samePrefixRandomID.AsByteArray(),
			zone:     12,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := TargetIsInZone(test.source, test.target, test.zone)
			assert.Equal(t, test.expected, actual)
		})
	}
}
