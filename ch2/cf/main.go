package main

import (
	"fmt"
	tempconv "gop1.io/ch2/tempconv0"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		float, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		celsius := tempconv.Celsius(float)
		fahrenheit := tempconv.Fahrenheit(float)
		fmt.Printf("%s = %s , %s = %s\n", celsius, tempconv.CToF(celsius), fahrenheit, tempconv.FToC(fahrenheit))
	}
}
