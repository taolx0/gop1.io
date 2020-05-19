package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //和C一样，隐式默认为空值
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] //s=s+sep+os.Arg[i]
		sep = " "
	}
	fmt.Println(s)
}
