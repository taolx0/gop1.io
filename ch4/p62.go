package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "¥"} //(看错代码。。)书上的代码运行不起来，应该是我的问题
	fmt.Println(RMB, symbol[RMB])
}
