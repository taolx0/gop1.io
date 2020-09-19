package main

//dup2 print count and text
//can read from files
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	//os.Args = []string{"dog.txt", "cat.txt"}
	//os.Args = []string{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count:%d\tinput.text():%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//注意：忽略input.Err()中可能的错误
}
