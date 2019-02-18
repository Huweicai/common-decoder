package decoder

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type UnixTimeStampDecoder struct {
}

func (d *UnixTimeStampDecoder) Sniffer(text string) Possibility {
	text = commonTimestampPreHanle(text)
	_, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return Impossible
	}
	//the recent years timestamp are 10 digits
	if len(text) != 10 {
		return AlmostImpossible
	}
	if strings.HasPrefix(text, "15") {
		return MayBe
	}
	return NotSure
}
func (d *UnixTimeStampDecoder) Decode(text string) (result string, ok bool) {
	text = commonTimestampPreHanle(text)
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		log.Println(text, err.Error())
		return
	}
	t := time.Unix(i, 0)
	return t.Format("2006-01-02 15:04:05"), true
}

func (d *UnixTimeStampDecoder) Encode(text string) (result string, ok bool) {
	return "", false
}

func commonTimestampPreHanle(text string) string {
	//timestamp may be divided by "," , such as 1,550,468,552
	text = strings.Replace(text, ",", "", -1)
	return text
}
