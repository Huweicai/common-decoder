package decoder

import (
	"encoding/base64"
	"log"
	"strings"
)

type Base64Decoder struct {
}

func (d *Base64Decoder) Sniffer(text string) Possibility {
	if strings.HasSuffix(text, "=") {
		return MayBe
	}
	return NotSure
}
func (d *Base64Decoder) Decode(text string) (result string, ok bool) {
	r, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Println(text, err.Error())
		return
	}
	return string(r), true
}

func (d *Base64Decoder) Encode(text string) (result string, ok bool) {
	result = base64.StdEncoding.EncodeToString([]byte(text))
	return result, true
}
