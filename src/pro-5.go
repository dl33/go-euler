package main

import (
	"fmt"
)

func main() {
	var smallest int = 2520

	isMultiple := true

	var max int = 1

	for i := 1 ; i < 21; i ++ {
		max *= i
	}

	fmt.Println("Max")
	fmt.Println(max)

	for j := 2520; j < max; j++ {
		smallest = j
		isMultiple = true
		for x := 1 ; x < 21; x++ {
			var val = smallest
			if val % x != 0 {
				isMultiple = false
				break
			} else {
				val /= x
			}
		}

		if isMultiple {
			break
		} else {
			smallest++
		}
	}

	fmt.Println("Smallest")
	fmt.Println(smallest)
}