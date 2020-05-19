package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fTOC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fTOC(boilingF))
}

func fTOC(f float64) float64 { //func 函数名（参数名 参数类型） 返回列表
	return (f - 32) * 5 / 9
}
