package decoder

import (
	"regexp"
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

var defaultTime = "00000000000000"

func (d *DateTimeDecoder) Decode(text string) (result string, ok bool) {
	text = preserveNumbers(text)
	if len(text) < len(defaultTime) {
		text += defaultTime[len(text):]
	}
	got, err := time.ParseInLocation("20060102150405", text, time.Local)
	if err != nil {
		return "", false
	}

	return strconv.FormatInt(got.Unix(), 10), true
}

func (d *DateTimeDecoder) Encode(text string) (result string, ok bool) {
	return d.Decode(text)
}

func preserveNumbers(s string) string {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return ""
	}
	return reg.ReplaceAllString(s, "")
}
