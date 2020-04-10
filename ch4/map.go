package main

import "fmt"

func main() {
	ages := map[string]int{
		"Jack":  35,
		"Alice": 22,
		"Bob":   27,
		"Frank": 59,
	}
	fmt.Println(ages["Jack"], ages["Alice"])
	delete(ages, "Alice")
	fmt.Println(ages["Alice"], ages["Jack"])
	ages["Jack"]++
	fmt.Println(ages["Jack"])

	/*built-in function map default random output*/
	for name, ages := range ages {
		fmt.Printf("name:%s\tage:%d\n", name, ages)
	}

	/*key sort*/
	var names []string
	for name := range ages {
		names = append(names, name)
	}
}
