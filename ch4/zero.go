package main

import "fmt"

func main() {
	var a [32]byte = [32]byte{1, 2, 3, 4, 5, 6, 7}
	zero(&a)
	fmt.Println(a)
}

/*func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}*/

func zero(ptr *[32]byte) {
	*ptr = [32]byte{} //*ptr代表数组的值，相当于全0值赋值给数组a
}
