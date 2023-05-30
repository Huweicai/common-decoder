package decoder

import (
	"log"
	"strconv"
	"time"
)

type UnixTimeStampNanoDecoder struct {
}

func (d *UnixTimeStampNanoDecoder) Sniffer(text string) Possibility {
	text = commonTimestampPreHanle(text)
	_, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return Impossible
	}

	if len(text) != 19 {
		return AlmostImpossible
	}

	return NotSure
}
func (d *UnixTimeStampNanoDecoder) Decode(text string) (result interface{}, ok bool) {
	text = commonTimestampPreHanle(text)
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		log.Println(text, err.Error())
		return
	}

	t := time.Unix(i/100_0000_000, 0)
	return t.Format("2006-01-02 15:04:05"), true
}

func (d *UnixTimeStampNanoDecoder) Encode(text string) (result interface{}, ok bool) {
	return "", false
}
