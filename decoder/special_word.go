package decoder

import (
	"strconv"
	"strings"
	"time"
)

type SpecialWordDecoder struct{}

func (d *SpecialWordDecoder) Sniffer(text string) Possibility {
	return NotSure
}

func buildSpecialWordResults(t time.Time) []*DecodeResult {
	return []*DecodeResult{
		{Result: strconv.FormatInt(t.Unix(), 10)},
		{Result: strconv.FormatInt(t.UnixNano(), 10)},
		{Result: t.Format("2006-01-02 15:04:05")},
	}
}

func (d *SpecialWordDecoder) Decode(text string) (result interface{}, ok bool) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "now":
		return buildSpecialWordResults(time.Now()), true
	case "today":
		today := time.Now().Truncate(24 * time.Hour)
		return buildSpecialWordResults(today), true
	case "yesterday":
		yesterday := time.Now().Add(-24 * time.Hour).Truncate(24 * time.Hour)
		return buildSpecialWordResults(yesterday), true
	case "tomorrow", "tommorow":
		tomorrow := time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour)
		return buildSpecialWordResults(tomorrow), true
	default:
		return "", false
	}
}

func (d *SpecialWordDecoder) Encode(text string) (result interface{}, ok bool) {
	return d.Decode(text)
}
