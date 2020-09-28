package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	//AbsoluteZeroF Fahrenheit = -273.15
	//FreezingF Fahrenheit = 32.0
	//BoilingF  Fahrenheit = 212.0
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
