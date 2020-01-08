package decoder

import "sort"

type Possibility int

const (
	Impossible = iota
	AlmostImpossible
	MayNotBe
	NotSure
	MayBe
	MustBe
)

var Decoders = map[string]Decoder{
	"Unix timestamp":      new(UnixTimeStampDecoder),
	"Unix timestamp mill": new(UnixTimeStampMillDecoder),
	"Unix timestamp nano": new(UnixTimeStampNanoDecoder),
	"Base64":              new(Base64Decoder),
	"URL decode":          new(URLDecoder),
	"Unicode":             new(UnicodeDecoder),
	"Date Time":           new(DateTimeDecoder),
	"Special Word":        new(SpecialWordDecoder),
}

type Decoder interface {
	Sniffer(text string) Possibility
	Decode(text string) (string, bool)
	Encode(text string) (string, bool)
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
		if level <= AlmostImpossible {
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

func Encode(text string) (results *DecodeResults) {
	results = &DecodeResults{}
	for name, dcd := range Decoders {
		result, ok := dcd.Encode(text)
		if !ok || text == result || result == "" {
			continue
		}
		results.Add(MayBe, name, result)
	}
	sort.Sort(results)
	return
}
