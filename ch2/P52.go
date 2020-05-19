package main

import "fmt"

func main() {
	s := "にほんご "
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
}
