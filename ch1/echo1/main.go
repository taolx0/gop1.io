package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //和C一样，隐式默认为空值
	os.Args = []string{"dog", "cat", "tiger"}
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
