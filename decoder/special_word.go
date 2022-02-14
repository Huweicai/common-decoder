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
		now := time.Now()
		return []*DecodeResult{
			{
				Result: strconv.FormatInt(now.Unix(), 10),
			},
			{
				Result: now.Format("2006-01-02 15:04:05"),
			},
		}, true
	default:
		return "", false
	}
}

func (d *SpecialWordDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
