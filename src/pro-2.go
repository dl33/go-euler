package main

import (
	"fmt"

)

func main() {
	var fibArray [1000000]int
	fibArray[0] = 1
	fibArray[1] = 2

	var sum int = fibArray[1]
	var i = 2
	for sum <= 4000000 {
		fibArray[i] = fibArray[i-1] + fibArray[i-2]

		if(fibArray[i] % 2 == 0) {
			sum += fibArray[i]
		}

		i++
	}

	fmt.Println(sum)
}