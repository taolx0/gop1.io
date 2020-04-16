package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	//AbsoluteZeroF Fahrenheit = -273.15
	FreezingF Fahrenheit = 32.0
	BoilingF  Fahrenheit = 212.0
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g°C\n", BoilingC-FreezingC)
	fmt.Printf("%g°F\n", BoilingF-CToF(FreezingC))
}
