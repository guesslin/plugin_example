package types // github.com/guesslin/plug/types

type Adder interface {
	Add(op1, op2 int) int
}

type Factory func() Adder
