package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", " "
	os.Args = []string{"dog", "cat", "lion"}
	for _, s2 := range os.Args[:] {
		s += s2 + seq
	}
	fmt.Println(s)
}
