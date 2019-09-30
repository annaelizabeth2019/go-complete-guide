package main

import "fmt"

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range ints {
		result := ""
		if i%2 == 0 {
			result = "even"
		} else {
			result = "odd"
		}
		fmt.Println(i, " is ", result)
	}
}
