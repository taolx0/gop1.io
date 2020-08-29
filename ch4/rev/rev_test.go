package rev

import (
	"fmt"
	"testing"
)

func TestRev(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //加上"..."需要划分成a[:]
	reverse(a[:])
	fmt.Println(a)
	/*左移n个元素的方法，此处n=2*/
	fmt.Println("左移n位元素，n=2")
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)
}
