package main

import "fmt"

func main() {
	data := []string{"one", "two", "", "four"}
	data = nonempty(data)
	fmt.Println(data)
}

func nonempty(strings []string) []string { //normal
	i := 0
	for _, i2 := range strings {
		if i2 != "" {
			strings[i] = i2
			i++
		}
	}
	return strings[:i]
}

/*func nonempty(strings []string) []string {	//test->error
	for i, i2 := range strings {
		if strings[i] != "" {
			strings[i] = i2
		}
	}
	return strings[:]
}*/
