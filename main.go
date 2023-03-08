package main

import (
	"fmt"
)

func main() {

	slice1 := []int{1, 2, 3}

	slice2 := make([]int, 0)

	for i, v := range slice1 {
		if v = 1 {
			slice2 = append(slice2, v)

		}
	}
	fmt.Println(slice2)
}
