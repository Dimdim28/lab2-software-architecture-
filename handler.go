package lab2

import (
	"io"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	bufRead, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := string(bufRead)
	trimmed := strings.Trim(expression, " \n")
	res, err := PrefixToInfix(trimmed)
	if err != nil {
		return err
	}
	ch.Output.Write([]byte(res))
	return nil
}
