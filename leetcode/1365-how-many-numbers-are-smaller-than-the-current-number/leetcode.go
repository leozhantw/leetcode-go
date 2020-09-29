package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	smallerNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i] > nums[j] {
				smallerNums[i]++
			}
		}
	}
	return smallerNums
}
