package decoder

import (
	"strconv"
	"strings"
	"time"
)

type DateTimeDecoder struct {
}

func (d *DateTimeDecoder) Sniffer(text string) Possibility {
	if !strings.HasPrefix(text, "2") {
		return AlmostImpossible
	}

	if !strings.HasPrefix(text, "20") {
		return MayNotBe
	}

	return NotSure
}

var defaultTime = "0000-00-00 00:00:00"

func (d *DateTimeDecoder) Decode(text string) (result string, ok bool) {
	if len(text) < len(defaultTime) {
		text += defaultTime[len(text):]
	}
	got, err := time.ParseInLocation("2006-01-02 15:04:05", text, time.Local)
	if err != nil {
		return "", false
	}

	return strconv.FormatInt(got.Unix(), 10), true
}

func (d *DateTimeDecoder) Encode(text string) (result string, ok bool) {
	return d.Decode(text)
}
