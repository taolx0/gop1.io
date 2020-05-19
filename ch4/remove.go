package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(remove(data, 2))
}

/*function one*/
/*func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}*/

/*function two*/
func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
