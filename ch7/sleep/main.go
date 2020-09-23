package main

import (
	"flag"
	"fmt"
	"time"
)

var duration = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Println(duration)
	time.Sleep(*duration)
	fmt.Println()
}
