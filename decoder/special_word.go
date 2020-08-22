package decoder

import (
	"strconv"
	"time"
)

type SpecialWordDecoder struct {
}

func (d *SpecialWordDecoder) Sniffer(text string) Possibility {
	return NotSure
}

func (d *SpecialWordDecoder) Decode(text string) (result interface{}, ok bool) {
	switch text {
	case "now":
		return strconv.FormatInt(time.Now().Unix(), 10), true
	default:
		return "", false
	}
}

func (d *SpecialWordDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
