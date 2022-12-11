package decoder

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

type UnitDecoder struct {
}

func (d *UnitDecoder) Sniffer(text string) Possibility {
	if strings.HasPrefix(text, "\\u") {
		return MayBe
	}
	return NotSure
}

var unitRegex = regexp.MustCompile(`(\d+\.*\d*)(\S{1,6})`)

type unitGroup struct {
	name             string
	unitCoefficients map[string]float64
	showUnits        []string
}

func NewUnitGroups() []*unitGroup {
	return []*unitGroup{
		{
			name: "STORAGE",
			unitCoefficients: map[string]float64{
				"b": 1,
				"B": 8,

				"KB": 1000 * 8,
				"MB": 1000 * 1000 * 8,
				"GB": 1000 * 1000 * 1000 * 8,
				"TB": 1000 * 1000 * 1000 * 1000 * 8,
				"PB": 1000 * 1000 * 1000 * 1000 * 1000 * 8,
				"EB": 1000 * 1000 * 1000 * 1000 * 1000 * 1000 * 8,

				"KiB": 1024 * 8,
				"MiB": 1024 * 1024 * 8,
				"GiB": 1024 * 1024 * 1024 * 8,
				"TiB": 1024 * 1024 * 1024 * 1024 * 8,
				"PiB": 1024 * 1024 * 1024 * 1024 * 1024 * 8,
				"EiB": 1024 * 1024 * 1024 * 1024 * 1024 * 1024 * 8,
			},
			showUnits: []string{"KB", "MB", "GiB", "GB", "TB", "TiB", "PB"},
		},
		{
			name: "TIME",
			unitCoefficients: map[string]float64{
				"ns": float64(time.Nanosecond),
				"us": float64(time.Microsecond),
				"μs": float64(time.Microsecond),

				"ms":   float64(time.Millisecond),
				"s":    float64(time.Second),
				"min":  float64(time.Minute),
				"h":    float64(time.Hour),
				"day":  float64(24 * time.Hour),
				"year": float64(365 * 24 * time.Hour),
			},
			showUnits: []string{"ms", "s", "min", "h", "day", "year"},
		},
		{
			name: "MONEY",
			unitCoefficients: map[string]float64{
				"￥":   1,
				"万￥":  10_000,
				"十万￥": 100_000,
				"百万￥": 1_000_000,
				"千万￥": 10_000_000,
				"亿￥":  100_000_000,
				"十亿￥": 1_000_000_000,
				"百亿￥": 10_000_000_000,
				"千亿￥": 100_000_000_000,
				"万亿￥": 1_000_000_000_000,
				"$":   dollarToRMBate * 1,
				"万$":  dollarToRMBate * 10_000,
				"十万$": dollarToRMBate * 100_000,
				"百万$": dollarToRMBate * 1_000_000,
				"千万$": dollarToRMBate * 10_000_000,
				"亿$":  dollarToRMBate * 100_000_000,
				"十亿$": dollarToRMBate * 1_000_000_000,
				"百亿$": dollarToRMBate * 10_000_000_000,
				"千亿$": dollarToRMBate * 100_000_000_000,
				"万亿$": dollarToRMBate * 1_000_000_000_000,
			},
			showUnits: []string{"￥", "万￥", "亿￥", "万亿￥", "亿$", "万亿$"},
		},
	}
}

// approximately 7
const dollarToRMBate = 7

func (d *UnitDecoder) Decode(text string) (interface{}, bool) {
	ret := unitRegex.FindStringSubmatch(text)
	if len(ret) != 3 {
		return nil, false
	}

	number, err := strconv.ParseFloat(ret[1], 64)
	if err != nil {
		return nil, false
	}
	unit := ret[2]

	groups := NewUnitGroups()

	for _, group := range groups {
		coefficient, ok := group.unitCoefficients[unit]
		if !ok {
			continue
		}

		baseNumber := number * coefficient

		var ret []*DecodeResult
		lo.ForEach(group.showUnits, func(showUnit string, _ int) {
			if showUnit == unit {
				return
			}

			showCoefficient := group.unitCoefficients[showUnit]
			number := baseNumber / showCoefficient

			// too small number is meaningless to show
			if number < 0.01 {
				return
			}
			ret = append(ret, &DecodeResult{
				Possibility: MustBe,
				Result:      fmt.Sprintf("%.2f%s", number, showUnit),
			})
		})

		return ret, true
	}

	return nil, false
}

func (d *UnitDecoder) Encode(text string) (result interface{}, ok bool) {
	textQuoted := strconv.QuoteToASCII(text)
	result = textQuoted[1 : len(textQuoted)-1]
	return result, true
}
