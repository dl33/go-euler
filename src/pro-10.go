package main

import (
	"fmt"
	"util"
)

func main()  {
	var sum int = 0

	for i := 2 ; i < 2000001 ; i++{
		if util.IsPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
