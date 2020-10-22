package leetcode

func permute(nums []int) [][]int {
	deep := len(nums)
	used, p, res := make([]bool, deep), []int{}, [][]int{}
	genPermutation(nums, 0, deep, &used, &p, &res)
	return res
}

func genPermutation(nums []int, index, deep int, used *[]bool, p *[]int, res *[][]int) {
	if index == deep {
		temp := make([]int, deep)
		copy(temp, *p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < deep; i++ {
		if !(*used)[i] {
			(*used)[i] = true
			*p = append(*p, nums[i])
			genPermutation(nums, index+1, deep, used, p, res)
			*p = (*p)[:len(*p)-1]
			(*used)[i] = false
		}
	}
	return
}
