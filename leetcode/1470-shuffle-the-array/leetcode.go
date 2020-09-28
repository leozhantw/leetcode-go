package leetcode

func shuffle(nums []int, n int) []int {
	shuffle := make([]int, n*2)

	for i := 0; i < n; i++ {
		shuffle[i*2] = nums[i]
		shuffle[(i*2)+1] = nums[n+i]
	}

	return shuffle
}
