package decoder

import (
	"fmt"
	"strconv"
)

type DecimalDecoder struct {
}

func (h DecimalDecoder) Sniffer(text string) Possibility {
	_, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return Impossible
	}

	return MustBe
}

func (h DecimalDecoder) Decode(text string) (interface{}, bool) {
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return nil, false
	}

	results := []*DecodeResult{
		{
			DecoderName: "Decimal to Binary",
			Result:      fmt.Sprintf("%b", i),
		},
		{
			DecoderName: "Decimal to Hex",
			Result:      fmt.Sprintf("%x", i),
		},
	}

	if 0 <= i && i <= 126 {
		results = append(results, &DecodeResult{
			DecoderName: "Decimal to ASCII",
			Result:      fmt.Sprintf("%c", i),
		})
	}

	return results, true
}

func (h DecimalDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
