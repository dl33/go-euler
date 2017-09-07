package util

import "math"

func IsPrime(num int) bool {
	var ret bool = true
	if num == 2 {
		return ret
	}

	if num % 2 != 0 {
		var ubound int = int(math.Sqrt(float64(num))) + 1

		for i := 2; i < ubound; i++ {
			if num % i == 0 {
				ret = false
				break
			}
		}
	} else {
		ret = false;
	}

	return ret;
}
