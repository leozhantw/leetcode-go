package leetcode

func xorOperation(n int, start int) int {
	var nums int
	for i := 0; i < n; i++ {
		nums ^= start + (2 * i)
	}
	return nums
}
