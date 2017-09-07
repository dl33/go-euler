package main

import (
	"fmt"
	"util"
)

func main()  {
	var sum int = 0

	for i := 1 ; i < 2000001 ; i++{
		if util.isPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
