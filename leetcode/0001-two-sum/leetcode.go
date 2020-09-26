package leetcode

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int, len(nums))

	for i, v := range nums {
		if _, ok := diffs[v]; ok {
			return []int{diffs[v], i}
		}

		diffs[target-v] = i
	}

	return nil
}
