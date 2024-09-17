package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) (int, error)
}

type CLIHandler struct {
	output     io.Writer
	calculator Calculator
}

func NewCLIHandler(calc Calculator, output io.Writer) *CLIHandler {
	return &CLIHandler{output, calc}
}

func (this *CLIHandler) Handle(inputs []string) error {
	if len(inputs) != 2 {
		return fmt.Errorf("%w: usage: go run main.go <a> <b>", errInvalidCountOfArgs)
	}
	addend1, err := strconv.Atoi(inputs[0])
	if err != nil {
		return fmt.Errorf("%w: %s", errInvalidArg, err)
	}
	addend2, err := strconv.Atoi(inputs[1])
	if err != nil {
		return fmt.Errorf("%w: %s", errInvalidArg, err)
	}
	result, err := this.calculator.Calculate(addend1, addend2)
	if err != nil {
		return fmt.Errorf("%w: %w", errCalculate, err)
	}

	if _, err := fmt.Fprintln(this.output, result); err != nil {
		return fmt.Errorf("%w: %w", errOutput, err)
	}
	return nil
}

var (
	errInvalidCountOfArgs = errors.New("too few arguments")
	errInvalidArg         = errors.New("invalid argument")
	errOutput             = errors.New("error output")
	errCalculate          = errors.New("error calculate")
)
