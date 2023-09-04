package decoder

import (
	"strings"

	"github.com/ianlancetaylor/demangle"
)

type DemangleDecoder struct {
}

func (h DemangleDecoder) Sniffer(text string) Possibility {
	return MayBe
}

func (h DemangleDecoder) Decode(text string) (interface{}, bool) {
	output, err := demangle.ToString(strings.TrimSpace(text))
	if err != nil {
		return nil, false
	}

	return &DecodeResult{
		DecoderName: "C++/Rust Symbol Demangle",
		Result:      output,
	}, true
}

func (h DemangleDecoder) Encode(text string) (interface{}, bool) {
	return nil, false
}
