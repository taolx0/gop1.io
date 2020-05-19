package main

import "fmt"

func main() {
	p := []string{"H", "E", "L", "L", "O"} //add "..." or not add "..." ,difference
	q := []string{"H", "E", "R", "L", "D"}
	fmt.Println(equal(p, q))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x { //function range's details
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
