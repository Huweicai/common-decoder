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

func (d *Base64Decoder) Decode(text string) (result interface{}, ok bool) {
	r, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		log.Println(text, err.Error())
		return
	}

	return string(r), true
}

func (d *Base64Decoder) Encode(text string) (result interface{}, ok bool) {
	result = base64.StdEncoding.EncodeToString([]byte(text))
	return result, true
}

type Base64URLDecoder struct {
}

func (d *Base64URLDecoder) Sniffer(text string) Possibility {
	return NotSure
}

func (d *Base64URLDecoder) Decode(text string) (result interface{}, ok bool) {
	r, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		log.Println(text, err.Error())
		return
	}

	return string(r), true
}

func (d *Base64URLDecoder) Encode(text string) (result interface{}, ok bool) {
	result = base64.RawURLEncoding.EncodeToString([]byte(text))
	return result, true
}
