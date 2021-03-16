package decoder

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
)

type JWTDecoder struct {
}

func (d *JWTDecoder) Sniffer(text string) Possibility {
	s := strings.Split(text, ".")
	if len(s) != 3 {
		return Impossible
	}

	return MayBe
}

type JWTToken struct {
	Header    json.RawMessage `json:"Header"`
	PayLoad   json.RawMessage `json:"PayLoad"`
	Signature string          `json:"Signature"`
}

func (d *JWTDecoder) Decode(text string) (result interface{}, ok bool) {
	log.Println(text, "JWT INPUT")
	s := strings.Split(text, ".")
	if len(s) != 3 {
		return nil, false
	}

	header, err := base64.RawURLEncoding.DecodeString(s[0])
	if err != nil {
		return nil, false
	}

	payload, err := base64.RawURLEncoding.DecodeString(s[1])
	if err != nil {
		return nil, false
	}

	return &JWTToken{
		Header:    header,
		PayLoad:   payload,
		Signature: s[2],
	}, true
}

func (d *JWTDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
