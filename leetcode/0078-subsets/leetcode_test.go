package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type give struct {
	nums []int
}

func Test_subsets(t *testing.T) {
	tests := []struct {
		give give
		want [][]int
	}{
		{
			give: give{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				nil,
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			give: give{
				nums: []int{0},
			},
			want: [][]int{
				nil,
				{0},
			},
		},
	}

	for _, tt := range tests {
		got := subsets(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
