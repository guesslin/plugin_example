package main

import (
	"flag"
	"fmt"
	"plugin"

	"github.com/guesslin/plug/types"
)

func main() {
	mode := flag.String("mode", "adder", "load which kind [adder|greet] of plugin")
	path := flag.String("path", "", "path to plugin file")
	flag.Parse()
	if *path == "" {
		fmt.Println("No path given")
		return
	}
	fmt.Println("======= main =======")
	switch *mode {
	case "adder":
		adder := loadAdder(*path)
		op1 := 100
		op2 := 20
		fmt.Printf("Adder %d + %d = %d result\n", op1, op2, adder.Add(op1, op2))
	case "greet":
		loadGreet(*path)
	default:
		fmt.Printf("mode %s not supported!\n", *mode)
	}
}

func loadPlugin(path string) *plugin.Plugin {
	p, err := plugin.Open(path)
	mustOK(err)
	return p
}

func loadAdder(path string) types.Adder {
	p := loadPlugin(path)
	f, err := p.Lookup("Factory")
	mustOK(err)
	factory, ok := f.(*types.Factory)
	if !ok {
		panic("assert failed")
	}
	return (*factory)()
}

func loadGreet(path string) {
	p := loadPlugin(path)
	gFunc, err := p.Lookup("Greet")
	mustOK(err)
	gFunc.(func(n string))("guesslin")
}

func mustOK(err error) {
	if err != nil {
		panic(err)
	}
}
