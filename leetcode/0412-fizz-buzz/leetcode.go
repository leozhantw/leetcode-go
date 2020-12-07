package leetcode

import "strconv"

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		var s string
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else {
			s = strconv.Itoa(i)
		}
		res = append(res, s)
	}
	return res
}
