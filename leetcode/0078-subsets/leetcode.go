package leetcode

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, num := range nums {
		for _, v := range res {
			clone := make([]int, len(v), len(v)+1)
			copy(clone, v)
			clone = append(clone, num)
			res = append(res, clone)
		}
	}
	return res
}
