package main

import "fmt"

func main() {
	for i, r := range "Hello,World！" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r) //%q,带引号的字符串
	}
}
