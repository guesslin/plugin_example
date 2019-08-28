package main

import (
	"github.com/guesslin/plug/plugs/adder/mod/alpha"
	"github.com/guesslin/plug/types"
)

type mod struct{}

func (mod) Add(op1, op2 int) int {
	return (op1 + op2) * int(10*alpha.Alpha)
}

var (
	Factory types.Factory
)

func init() {
	Factory = func() types.Adder { return &mod{} }
}
