package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5)
	x = append(x, 6)
	x = append(x, x...)
	fmt.Println(x)
}
