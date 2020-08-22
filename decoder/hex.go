package decoder

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type HexDecoder struct {
}

func (h HexDecoder) Sniffer(text string) Possibility {
	if !strings.HasPrefix(text, "0x") {
		return Impossible
	}

	return MayBe
}

func (h HexDecoder) Decode(text string) (interface{}, bool) {
	text = strings.TrimPrefix(text, "0x")
	tmp, err := hex.DecodeString(text)
	if err != nil {
		return nil, false
	}

	binary := strings.TrimPrefix(fmt.Sprintf("%.8b", tmp), "[")
	binary = strings.TrimSuffix(binary, "]")
	ret := []*DecodeResult{
		{
			DecoderName: "Hex to Binary",
			Result:      binary,
		}}

	decimal, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		return ret, true
	}

	return append(ret, &DecodeResult{
		DecoderName: "Hex to Decimal",
		Result:      decimal,
	}), true
}

func (h HexDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
