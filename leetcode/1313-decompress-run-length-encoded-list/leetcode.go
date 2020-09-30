package leetcode

func decompressRLElist(nums []int) []int {
	var list []int
	for i := 0; i < len(nums)/2; i++ {
		freq := nums[2*i]
		val := nums[2*i+1]
		for j := 0; j < freq; j++ {
			list = append(list, val)
		}
	}
	return list
}
