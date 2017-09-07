package main

import (
	"fmt"
	"math"
)

func main()  {
	var no int = 0
	var num int = 2

	for {
		if isPrime(num) {
			no ++

			if no == 10001 {
				break
			}
		}
		num++
	}

	fmt.Println(num)
}
