package main

import "fmt"

func main() {
	/*s=[]int{}
	fmt.Println(len(s),s)*/

	/*function append*/
	var runes []rune
	for _, i2 := range "Hello,World!" { //add string's ASCII , not 0,1,2,3...
		runes = append(runes, i2)
	}
	fmt.Printf("%q ", runes)
}
