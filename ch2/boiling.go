package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g° or %g°C\n", f, c)
	//boiling = 212°F or 100°C
}
