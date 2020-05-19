package main

import "fmt"

//const noDelay time.Duration = 0
//const timout = 5 * time.Minute

const (
	a = 1
	b
	c = 2
	d
)

func main() {
	/*fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timout)
	fmt.Printf("%T %[1]v\n", time.Minute)*/
	fmt.Println(a, b, c, d)
}
