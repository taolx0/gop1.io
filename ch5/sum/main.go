package main

import "fmt"

/*variable-length function*/
func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 9))
	fmt.Println(sum(1, 5, 6, 7))
	//当实参已经存在于slice,how to call a function
	fmt.Println(sum(values...))
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
}

func sum(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func f(...int) {}
func g([]int)  {}
