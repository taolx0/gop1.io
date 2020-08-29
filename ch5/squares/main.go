package main

import (
	"fmt"
	"reflect"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println("typeof(f):", reflect.TypeOf(f))
	fmt.Println("typeof(f()):", reflect.TypeOf(f()))
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	//wrong,can not work
	//fmt.Println(squares()())
	//fmt.Println(squares()())
	//fmt.Println(squares()())
	//fmt.Println(squares()())
}
