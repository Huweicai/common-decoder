package decoder

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"net"
)

type CIDRDecoder struct {
}

func (d *CIDRDecoder) Sniffer(text string) Possibility {
	_, _, err := net.ParseCIDR(text)
	if err != nil {
		return Impossible
	}

	return MustBe
}

type CIDRToken struct {
	Header    json.RawMessage `json:"Header"`
	PayLoad   json.RawMessage `json:"PayLoad"`
	Signature string          `json:"Signature"`
}

func (d *CIDRDecoder) Decode(text string) (result interface{}, ok bool) {
	_, cidr, err := net.ParseCIDR(text)
	if err != nil {
		return nil, false
	}

	ones, total := cidr.Mask.Size()
	num := math.Pow(2, float64(total-ones))

	mask := binary.BigEndian.Uint32(cidr.Mask)
	start := binary.BigEndian.Uint32(cidr.IP)

	// find the final address
	last := (start & mask) | (mask ^ 0xffffffff)
	lastIP := make(net.IP, 4)
	binary.BigEndian.PutUint32(lastIP, last)

	return &DecodeResult{
		Possibility: MustBe,
		Result:      fmt.Sprintf("IP Count: %d Range: %s->%s", int(num), cidr.IP.String(), lastIP.String()),
	}, true
}

func (d *CIDRDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
