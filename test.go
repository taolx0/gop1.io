package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "hello"
	//s1 := make([]byte, 5, 10)
	slice := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s[:3], reflect.TypeOf(s[:3]))
	fmt.Println(string(slice[:3]), reflect.TypeOf(slice[:3]))

	for i := range s {
		fmt.Println(i)
	}
}
