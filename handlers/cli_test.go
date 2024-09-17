package handlers

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	"github.com/SmartyJerry/calc-lib/calc"
)

func assertEqual(t *testing.T, expected, actual any) {
	if !reflect.DeepEqual(expected, actual) {
		t.Helper()
		t.Errorf("\n"+
			"expected: [%v]\n"+
			"actual: [%v]", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error) {
	if !errors.Is(actual, expected) {
		t.Helper()
		t.Errorf("\n"+
			"expected: [%v]\n"+
			"actual: [%v]", expected, actual)
	}
}

func TestHandler_TooFewArgs(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{})
	err := handler.Handle(nil)
	assertError(t, errInvalidCountOfArgs, err)
}

func TestHandler_TooManyArgs(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{})
	err := handler.Handle([]string{"1", "2", "3"})
	assertError(t, errInvalidCountOfArgs, err)
}

func TestHandler_InvalidFirstArg(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{})
	err := handler.Handle([]string{"a", "1"})
	assertError(t, errInvalidArg, err)
}
func TestHandler_InvalidSecondArg(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{})
	err := handler.Handle([]string{"1", "12w"})
	assertError(t, errInvalidArg, err)
}
func TestHandler_InvalidOutput(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{err: errors.New("err")})
	err := handler.Handle([]string{"1", "2"})
	assertError(t, errOutput, err)
}
func TestHandler_ValidInputNoError(t *testing.T) {
	handler := NewCLIHandler(calc.Addition{}, fakeWriter{})
	err := handler.Handle([]string{"1", "2"})
	assertError(t, nil, err)
}

func TestHandler_CalculateError(t *testing.T) {
	handler := NewCLIHandler(calc.Division{}, fakeWriter{})
	err := handler.Handle([]string{"1", "0"})
	assertError(t, errCalculate, err)
}

type fakeWriter struct {
	bytes.Buffer
	err error
}

func (this fakeWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
