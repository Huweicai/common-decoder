package decoder

import (
	"encoding/json"
	"log"
	"reflect"
	"sort"
	"strings"
)

type Possibility int

const (
	Impossible = iota
	AlmostImpossible
	MayNotBe
	NotSure
	MayBe
	MustBe
)

var Decoders = []Decoder{
	new(UnixTimeStampDecoder),
	new(UnixTimeStampMillDecoder),
	new(UnixTimeStampNanoDecoder),
	new(Base64Decoder),
	new(Base64URLDecoder),
	new(URLDecoder),
	new(UnicodeDecoder),
	new(DateTimeDecoder),
	new(SpecialWordDecoder),
	new(HexDecoder),
	new(DecimalDecoder),
	new(JWTDecoder),
	new(CIDRDecoder),
}

type Decoder interface {
	Sniffer(text string) Possibility
	Decode(text string) (interface{}, bool) // string, []*DecodeResult, *DecodeResult are all acceptable
	Encode(text string) (interface{}, bool)
}

type DecodeResults struct {
	Data []*DecodeResult
}

func (d *DecodeResults) Less(i, j int) bool {
	return d.Data[i].Possibility < d.Data[j].Possibility
}

func (d *DecodeResults) Len() int {
	return len(d.Data)
}
func (d *DecodeResults) Swap(i, j int) {
	d.Data[i], d.Data[j] = d.Data[j], d.Data[i]
}

func (d *DecodeResults) Add(level Possibility, name string, result interface{}) {
	switch t := result.(type) {
	case []*DecodeResult:
		for _, v := range t {
			if v.Possibility == 0 {
				v.Possibility = level
			}

			// patch decoder name if not exist
			if v.DecoderName == "" {
				v.DecoderName = name
			}
		}
		d.Data = append(d.Data, t...)

	case *DecodeResult:
		// patch decoder name if not exist
		if t.DecoderName == "" {
			t.DecoderName = name
		}
		if t.Possibility == 0 {
			t.Possibility = level
		}

		d.Data = append(d.Data, t)
	case string:
		d.Data = append(d.Data, &DecodeResult{
			level,
			name,
			t})
	default:
		ret, err := json.Marshal(t)
		if err != nil {
			log.Println("unexpected result", err, t)
			break
		}
		d.Data = append(d.Data, &DecodeResult{
			level,
			name,
			string(ret)})
	}
}

type Namer interface {
	Name() string
}

type DecodeResult struct {
	Possibility Possibility
	DecoderName string
	Result      interface{}
}

func GetDecoderName(i interface{}) string {
	if namer, ok := i.(Namer); ok {
		return namer.Name()
	}

	name := reflect.TypeOf(i).Elem().Name()
	return strings.TrimSuffix(name, "Decoder")
}

func Decode(text string) (results *DecodeResults) {
	results = &DecodeResults{}

	for _, decoder := range Decoders {
		level := decoder.Sniffer(text)
		if level <= AlmostImpossible {
			continue
		}

		result, ok := decoder.Decode(text)
		if !ok || text == result {
			continue
		}

		name := GetDecoderName(decoder)

		results.Add(level, name, result)
	}

	sort.Sort(results)
	return
}

func Encode(text string) (results *DecodeResults) {
	results = &DecodeResults{}
	for _, decoder := range Decoders {
		result, ok := decoder.Encode(text)
		if !ok || text == result || result == "" {
			continue
		}

		name := GetDecoderName(decoder)

		results.Add(MayBe, name, result)
	}

	sort.Sort(results)
	return
}
