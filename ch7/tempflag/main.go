package main

import (
	"flag"
	"fmt"
	"gop1.io/ch7/tempcov"
)

var temp = tempcov.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
