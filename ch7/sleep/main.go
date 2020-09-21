package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	duration := flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Println(duration)
	time.Sleep(*duration)
}
