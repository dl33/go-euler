package main

import (
	"fmt"
	"strconv"
)

func main() {
	var largest int = 0

	for x := 0 ; x < 1000; x++ {
		for y := 0 ; y < 1000; y++ {
			val := x * y

			if isPalindromic(val) {

				if(largest < val) {
					largest = val
				}
			}
		}
	}

	fmt.Println("Largest")
	fmt.Println(largest)
}

func isPalindromic(num int) bool {
	var numStr string = strconv.Itoa(num)
	var numLen = len(numStr)

	var ret = true;
	if (numLen % 2 == 0) {
		for i := 0; i < numLen / 2; i++ {
			if (numStr[i] != numStr[numLen - i - 1]) {
				ret = false
				break
			}
		}
	} else {
		for i := 0; i < numLen / 2 + 1; i++ {
			if (numStr[i] != numStr[numLen - i - 1]) {
				ret = false
				break
			}
		}
	}

	return ret
}