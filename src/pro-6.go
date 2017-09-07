package main

import (
	"fmt"
)

func main()  {
	var squaresVal int = 0
	var sumVal int = 0

	for i:=1; i < 101; i++ {
		squaresVal += i * i
		sumVal += i
	}

	difVal := sumVal * sumVal - squaresVal

	fmt.Println(difVal)
}

