package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/Dimdim28/lab2-software-architecture-"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFromFile   = flag.String("f", "", "Take input from file")
	outputToFile    = flag.String("o", "", "Save outut to file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFromFile != "" {
		f, err := os.Open(*inputFromFile)
		if err != nil {
			os.Stderr.WriteString("Error opening file")
		}
		defer f.Close()
		input = f
	}

	if *outputToFile != "" {
		file, err := os.Create(*outputToFile)
		if err != nil {
			os.Stderr.WriteString("Error")
		}
		defer file.Close()
		output = file
	}

	if input == nil {
		os.Stderr.WriteString("No input provided")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if err != nil {
		println(err)
	}
}
