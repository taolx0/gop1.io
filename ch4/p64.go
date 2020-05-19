package main

import "fmt"

func main() {
	months := [...]string{1: "january", 2: "February", 3: "March", 4: "April", 5: "May",
		6: "June", 7: "July", 8: "August", 9: "September", 10: "October",
		11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer { //数组range迭代循环参数是数组，s依次表示数组的值，
		for _, q := range Q2 { // 指针指向数组，依次表示0，1，2，3，这个理解对吗？
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
}
