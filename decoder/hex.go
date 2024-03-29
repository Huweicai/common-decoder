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
	return MayBe
}

func (h HexDecoder) Decode(text string) (interface{}, bool) {
	text = strings.Replace(strings.TrimPrefix(text, "0x"), " ", "", -1)
	tmp, err := hex.DecodeString(text)
	if err != nil {
		return nil, false
	}

	binary := strings.TrimPrefix(fmt.Sprintf("%.8b", tmp), "[")
	binary = strings.TrimSuffix(binary, "]")
	results := []*DecodeResult{
		{
			DecoderName: "Hex to Binary",
			Result:      binary,
		}}

	decimal, err := strconv.ParseInt(text, 16, 64)
	if err != nil {
		return results, true
	}

	if 0 <= decimal && decimal <= 126 {
		results = append(results, &DecodeResult{
			DecoderName: "Hex to ASCII",
			Result:      fmt.Sprintf("%c", decimal),
		})
	}

	return append(results, &DecodeResult{
		DecoderName: "Hex to Decimal",
		Result:      decimal,
	}), true
}

func (h HexDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
