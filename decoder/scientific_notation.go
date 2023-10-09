package decoder

import (
	"strings"

	"github.com/shopspring/decimal"
)

type ScientificNotationDecoder struct {
}

func (h ScientificNotationDecoder) Sniffer(text string) Possibility {
	return MayBe
}

func (h ScientificNotationDecoder) Decode(text string) (interface{}, bool) {
	if !strings.Contains(text, "e") && !strings.Contains(text, "E") {
		return nil, false
	}

	num, err := decimal.NewFromString(strings.TrimSpace(text))
	if err != nil {
		return nil, false
	}

	return &DecodeResult{
		DecoderName: "ScientificNotation to Decimal",
		Result:      num.String(),
	}, true
}

func (h ScientificNotationDecoder) Encode(text string) (interface{}, bool) {
	return h.Decode(text)
}
