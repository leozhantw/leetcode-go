package leetcode

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		copy(target[index[i]+1:], target[index[i]:])
		target[index[i]] = nums[i]
	}
	return target
}
