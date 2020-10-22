package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_permute(t *testing.T) {
	tests := []struct {
		give give
		want [][]int
	}{
		{
			give: give{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		got := permute(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
