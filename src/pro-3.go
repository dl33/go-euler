package main

import (
	"fmt"
	"math"
)

func main() {
	//var num int = 600851475143
	var num int = 600851475143

	var largestFactor int
	for i := int(math.Sqrt(float64(num))); i > 0; i-- {
		var factor = i

		fmt.Println(factor);

		if (factor % 2 == 0) {
			continue
		} else if factor % 6 != 1 && factor % 6 != 5 {
			continue
		} else if num % factor != 0 && num != factor{
			continue
		}else {
			var ubound int = int(math.Sqrt(float64(factor)))
			var isFactor bool = true
			for j := 5; j < ubound; j+=6 {
				if factor % j == 0 || factor % (j + 2) == 0 {
					isFactor = false
					break
				}
			}

			if isFactor {
				largestFactor = factor
				break
			}
		}
	}

	fmt.Println("largestFactor")
	fmt.Println(largestFactor)
}