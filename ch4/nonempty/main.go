package main

import "fmt"

func main() {
	data := []string{"one", "two", "", "four"}
	data = nonempty(data)
	fmt.Println(data)
}

/*function one*/
/*func nonempty(strings []string) []string {
	i := 0
	for _, i2 := range strings {
		if i2 != "" {
			strings[i] = i2
			i++
		}
	}
	return strings[:i]
}*/

/*function two*/
func nonempty(strings []string) []string {
	out := strings[:0]
	for _, i2 := range strings {
		if i2 != "" {
			out = append(out, i2)
		}
	}
	return out
}

/*error example*/
/*func nonempty(strings []string) []string {
	for i, i2 := range strings {
		if strings[i] != "" {
			strings[i] = i2
		}
	}
	return strings[:]
}*/
