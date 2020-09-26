package leetcode

import (
	"fmt"
	"testing"
)

type give struct {
	nums   []int
	target int
}

func Test_twoSum(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			give: give{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
	}

	for _, t := range tests {
		got := twoSum(t.give.nums, t.give.target)
		fmt.Printf("【Your input】 %v 【Output】 %v\n", t.give, got)
	}
}
