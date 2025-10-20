package decoder

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type UnixTimeStampDecoder struct {
}

func (d *UnixTimeStampDecoder) parseUnixTimestamp(text string) (time.Time, string, bool) {
	text = commonTimestampPreHandle(text)
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		log.Println(text, err.Error())
		return time.Time{}, "", false
	}

	length := len(text)
	var t time.Time
	var decoderName string

	switch {
	case length <= 10:
		t = time.Unix(i, 0)
		decoderName = "Unix Timestamp"
	case length <= 13:
		t = time.Unix(0, i*int64(time.Millisecond))
		decoderName = "Unix Timestamp Milli"
	case length <= 16:
		t = time.Unix(0, i*int64(time.Microsecond))
		decoderName = "Unix Timestamp Micro"
	case length <= 19:
		t = time.Unix(0, i)
		decoderName = "Unix Timestamp Nano"
	default:
		log.Println(text, "unsupported timestamp length")
		return time.Time{}, "", false
	}

	year := t.Year()
	if year < 1970 || year > 3000 {
		return time.Time{}, "", false
	}

	return t, decoderName, true
}

func (d *UnixTimeStampDecoder) Sniffer(text string) Possibility {
	_, _, ok := d.parseUnixTimestamp(text)
	if !ok {
		return AlmostImpossible
	}
	return MayBe
}

func (d *UnixTimeStampDecoder) Decode(text string) (interface{}, bool) {
	t, decoderName, ok := d.parseUnixTimestamp(text)
	if !ok {
		return nil, false
	}

	baseTime := t.Format("2006-01-02 15:04:05")

	nano := t.Nanosecond()
	if nano > 0 {
		nanoStr := fmt.Sprintf("%09d", nano)
		fractionalPart := nanoStr[:3] + "ms," + nanoStr[3:6] + "us," + nanoStr[6:] + "ns"
		baseTime += " " + fractionalPart
	}

	return &DecodeResult{
		Possibility: MustBe,
		DecoderName: decoderName,
		Result:      baseTime,
	}, true
}

func (d *UnixTimeStampDecoder) Encode(text string) (interface{}, bool) {
	return nil, false
}

func commonTimestampPreHandle(text string) string {
	text = strings.TrimSpace(text)
	text = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, text)

	return text
}
