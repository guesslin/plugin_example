package main

import (
	"github.com/guesslin/plug/types"
)

type average struct{}

func (average) Add(op1, op2 int) int {
	return (op1 + op2) / 2
}

var (
	Factory types.Factory
)

func init() {
	Factory = func() types.Adder { return &average{} }
}
