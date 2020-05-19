package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //加上"..."需要划分成a[:]
	reverse(a[:])
	fmt.Println(a)
	/*左移n个元素的方法，此处n=2*/
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[j], s[i] = s[i], s[j]
	}
}
