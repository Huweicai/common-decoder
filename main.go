package main

import (
	"common-decoder/alfred"
	"common-decoder/decoder"
	"flag"
	"log"
)

func init()  {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) <= 1 {
		panic("too few arguments")
	}
	switch args[0] {
	case "decode":
		decode(args[1])
	case "encode":
	}
}

func decode(text string)   {
	results := decoder.Decode(text)
	otp := alfred.NewOutput()
	for _, result := range results.Data() {
		otp.AddSimpleTip(result.Result , result.DecoderName , "" , result.Result)
	}
	otp.Show()
}