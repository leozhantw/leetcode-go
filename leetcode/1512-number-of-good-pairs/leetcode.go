package leetcode

func numIdenticalPairs(nums []int) int {
	var pairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				pairs++
			}
		}
	}
	return pairs
}
