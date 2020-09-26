package leetcode

func runningSum(nums []int) []int {
	var sum int

	for i, v := range nums {
		nums[i] += sum
		sum += v
	}

	return nums
}
