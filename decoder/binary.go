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

	return []*DecodeResult{
		{
			DecoderName: "Binary to Decimal",
			Result:      decimal,
		},
		{
			DecoderName: "Binary to Hex",
			Result:      fmt.Sprintf("%x", decimal),
		},
	}, true
}

func (h BinaryDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
