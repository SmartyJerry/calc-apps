package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"calc-apps/handlers"

	"github.com/SmartyJerry/calc-lib/calc"
)

func main() {
	var (
		output io.Writer = os.Stdout
	)
	var op string
	flag.StringVar(&op, "op", "+", "One of + - * /")
	flag.Parse()
	calculator, ok := calculators[op]
	if !ok {
		fmt.Fprintln(os.Stderr, "Unknown operation")
	}
	handler := handlers.NewCLIHandler(calculator, output)

	err := handler.Handle(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
