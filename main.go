package main

import (
	"github.com/akamensky/argparse"
	"github.com/graniet/GoFuzz/fuzzer"
	"github.com/common-nighthawk/go-figure"
	"os"
	"fmt"
)

func main(){
	parser := argparse.NewParser("Gofuzz", "API Rest fuzzer written in Golang")
	file := parser.String("r", "requests", &argparse.Options{Required:true, Help:"File to API requests dump."})
	fuzzType := parser.String("t", "type", &argparse.Options{Required: false, Help: "Type of fuzzing: sql, xss, python"})
	printPayload := parser.Flag("v", "verbose", &argparse.Options{Required:false, Help:"Print payload verbose in checking process"})
	customPayload := parser.String("c", "custom", &argparse.Options{Required:false, Help:"Custom payload file"})
	formatRequest := parser.Flag("p", "postman", &argparse.Options{Required: false, Help:"Use postman format"})

	err := parser.Parse(os.Args)
	if err != nil{
		fmt.Print(parser.Usage("error: " + err.Error()))
		return
	}
	fmt.Printf("\n")
	figure.NewFigure("GoFuzz", "colossal", true).Print()
	fmt.Printf("\n")

	if *customPayload == "" && *fuzzType == ""{

		fmt.Print("error: Please set one type or custom payload path.")
		return
	}

	Flag := fuzzer.Flag{
		Verbose: *printPayload,
	}

	configuration := fuzzer.GoFuzz{
		Target: *file,
		ParamUsed: make(map[string]string),
		Fuzzer: fuzzer.Vulnerability{
			Type: *fuzzType,
		},
		Flags:Flag,
		CustomPayload: *customPayload,
		FormatPostMan: *formatRequest,
	}

	configuration.Run()
}
