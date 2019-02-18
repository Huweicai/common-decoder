package decoder

import (
	"log"
	"net/url"
)

type URLDecoder struct {
}

func (d *URLDecoder) Sniffer(text string) Possibility {
	return NotSure
}
func (d *URLDecoder) Decode(text string) (result string, ok bool) {
	result, err := url.QueryUnescape(text)
	if err != nil {
		log.Println(text, err.Error())
		return
	}
	return result, true
}

func (d *URLDecoder) Encode(text string) (result string, ok bool) {
	result = url.QueryEscape(text)
	return result, true
}
