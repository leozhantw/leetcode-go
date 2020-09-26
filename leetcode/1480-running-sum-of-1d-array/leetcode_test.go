package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_runningSum(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			give: give{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			give: give{
				nums: []int{3, 1, 2, 10, 1},
			},
			want: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tt := range tests {
		got := runningSum(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
