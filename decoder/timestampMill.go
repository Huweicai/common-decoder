package decoder

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type UnixTimeStampMillDecoder struct {
}

func (d *UnixTimeStampMillDecoder) Sniffer(text string) Possibility {
	text = commonTimestampPreHanle(text)
	_, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return Impossible
	}
	if len(text) != 13 {
		return AlmostImpossible
	}
	if strings.HasPrefix(text, "15") {
		return MayBe
	}
	return NotSure
}
func (d *UnixTimeStampMillDecoder) Decode(text string) (result string, ok bool) {
	text = commonTimestampPreHanle(text)
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		log.Println(text, err.Error())
		return
	}
	t := time.Unix(i/1000, 0)
	return t.Format("2006-01-02 15:04:05"), true
}

func (d *UnixTimeStampMillDecoder) Encode(text string) (result string, ok bool) {
	return "", false
}
