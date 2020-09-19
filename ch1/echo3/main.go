package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Args = []string{"dog", "cat", "lion"}
	fmt.Println(strings.Join(os.Args[:], " "))
}
