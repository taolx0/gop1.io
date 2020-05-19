package main

import (
	"crypto/sha256"
	"fmt"
)

func main() { //和书中结果不一致，是什么问题？
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}
