package main

type Flags int

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
)

func main() {

}

func IsUp(v Flags) {

}
