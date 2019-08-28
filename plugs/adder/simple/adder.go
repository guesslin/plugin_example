package main

import (
	"github.com/guesslin/plug/types"
)

type simple struct{}

func (s *simple) Add(op1, op2 int) int {
	return op1 + op2
}

var Factory types.Factory

func init() {
	Factory = func() types.Adder {
		return &simple{}
	}
}
