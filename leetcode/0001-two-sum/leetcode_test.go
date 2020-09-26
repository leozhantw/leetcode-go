package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		{
			give: give{
				nums:   []int{1, 2},
				target: 4,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.give.nums, tt.give.target)
		assert.Equal(t, tt.want, got)
	}
}
