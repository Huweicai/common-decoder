package main

import (
	"common-decoder/alfred"
	"common-decoder/decoder"
	"flag"
	"log"
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
	switch args[0] {
	case "decode":
		output = decode(args[1])
	case "encode":
		output = encode(args[1])
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
	for _, result := range results.Data() {
		otp.AddSimpleTip(result.Result, result.DecoderName, result.Result, "")
	}
	return otp
}

func encode(text string) *alfred.Output {
	results := decoder.Encode(text)
	otp := alfred.NewOutput()
	for _, result := range results.Data() {
		otp.AddSimpleTip(result.Result, result.DecoderName, result.Result, "")
	}
	return otp
}
