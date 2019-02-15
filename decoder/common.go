package decoder

import "sort"

type Possibility int

const (
	Impossible = iota
	MayNotBe
	NotSure
	MayBe
	MustBe
)

var Decoders = map[string]Decoder{
	"Unix timestamp":      new(UnixTimeStampDecoder),
	"Unix timestamp nano": new(UnixTimeStampMillDecoder),
	"Base64":              new(Base64Decoder),
	"URL decode":          new(URLDecoder),
}

type Decoder interface {
	Sniffer(text string) Possibility
	Decode(text string) (string, bool)
}

type DecodeResults struct {
	data []DecodeResult
}

func (d *DecodeResults) Less(i, j int) bool {
	return d.data[i].Possibility < d.data[j].Possibility
}

func (d *DecodeResults) Len() int {
	return len(d.data)
}
func (d *DecodeResults) Swap(i, j int) {
	d.data[i], d.data[j] = d.data[j], d.data[i]
}

func (d *DecodeResults) Data() []DecodeResult {
	return d.data
}

func (d *DecodeResults) Add(level Possibility, name, result string) {
	d.data = append(d.data, DecodeResult{level, name, result})
}

type DecodeResult struct {
	Possibility Possibility
	DecoderName string
	Result      string
}

func Decode(text string) (results *DecodeResults) {
	results = &DecodeResults{}
	for name, dcd := range Decoders {
		level := dcd.Sniffer(text)
		if level == Impossible {
			continue
		}
		result, ok := dcd.Decode(text)
		if !ok || text == result {
			continue
		}
		results.Add(level, name, result)
	}
	sort.Sort(results)
	return
}
