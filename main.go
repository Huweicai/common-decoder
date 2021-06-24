package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/Huweicai/common-decoder/alfred"
	"github.com/Huweicai/common-decoder/decoder"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 1 {
		panic("too few arguments")
	}

	var output *alfred.Output
	arg := strings.Join(args[1:], " ")
	switch args[0] {
	case "decode":
		output = decode(arg)
	case "encode":
		output = encode(arg)
	}

	if output == nil || len(output.Items) == 0 {
		output = alfred.NewOutput()
		output.AddSimpleTip("Nothing Found", "Nothing Found", "", "")
	}

	output.Show()
}

func decode(text string) *alfred.Output {
	results := decoder.Decode(text)
	otp := alfred.NewOutput()
	for _, result := range results.Data {
		str := fmt.Sprintf("%v", result.Result)

		otp.AddSimpleTip(str, result.DecoderName, str, "")
	}
	return otp
}

func encode(text string) *alfred.Output {
	results := decoder.Encode(text)
	otp := alfred.NewOutput()
	for _, result := range results.Data {
		str := fmt.Sprintf("%v", result.Result)

		otp.AddSimpleTip(str, result.DecoderName, str, "")
	}
	return otp
}
