package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//ignore input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count:%d\tinput.text():%s\n", n, line)
		}
	}
}
