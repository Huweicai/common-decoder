package decoder

import (
	"fmt"
	"strconv"
	"strings"
)

type BinaryDecoder struct {
}

func (h BinaryDecoder) Sniffer(text string) Possibility {
	return MayBe
}

func (h BinaryDecoder) Decode(text string) (interface{}, bool) {
	text = strings.Replace(strings.TrimPrefix(text, "0b"), " ", "", -1)
	decimal, err := strconv.ParseInt(text, 2, 64)
	if err != nil {
		return nil, false
	}

	results := []*DecodeResult{
		{
			DecoderName: "Binary to Decimal",
			Result:      decimal,
		},
		{
			DecoderName: "Binary to Hex",
			Result:      fmt.Sprintf("%x", decimal),
		},
	}

	if 0 <= decimal && decimal <= 126 {
		results = append(results, &DecodeResult{
			DecoderName: "Decimal to ASCII",
			Result:      fmt.Sprintf("%c", decimal),
		})
	}

	return results, true
}

func (h BinaryDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
