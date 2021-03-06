package decoder

import (
	"fmt"
	"strconv"
	"strings"
)

type UnicodeDecoder struct {
}

func (d *UnicodeDecoder) Sniffer(text string) Possibility {
	if strings.HasPrefix(text, "\\u") {
		return MayBe
	}
	return NotSure
}
func (d *UnicodeDecoder) Decode(text string) (interface{}, bool) {
	ret := ""
	sUnicodev := strings.Split(text, "\\u")
	for _, v := range sUnicodev {
		if len(v) < 1 {
			continue
		}

		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			return "", false
		}

		ret += fmt.Sprintf("%c", temp)
	}

	return ret, true
}

func (d *UnicodeDecoder) Encode(text string) (result interface{}, ok bool) {
	textQuoted := strconv.QuoteToASCII(text)
	result = textQuoted[1 : len(textQuoted)-1]
	return result, true
}
