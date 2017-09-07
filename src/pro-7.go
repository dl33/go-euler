package main

import (
	"fmt"
	"util"
)

func main()  {
	var no int = 0
	var num int = 2

	for {
		if util.IsPrime(num) {
			no ++

			if no == 10001 {
				break
			}
		}
		num++
	}

	fmt.Println(num)
}
