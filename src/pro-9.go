package main

import (
	"fmt"
)

func main()  {
	var a, b, c int
	for a = 1; a < 1001; a++ {
		for b = 1; b < 1001; b++ {
			for c = 1; c < 1001; c++ {
				if a + b + c == 1000 {
					if a * a + b * b == c * c {
						fmt.Println(a)
						fmt.Println(b)
						fmt.Println(c)
						return
					}
				}
			}
		}
	}
}
