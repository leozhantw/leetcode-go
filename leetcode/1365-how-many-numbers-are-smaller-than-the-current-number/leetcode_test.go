package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_smallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums: []int{8, 1, 2, 2, 3},
			},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			give: give{
				nums: []int{6, 5, 4, 8},
			},
			want: []int{2, 1, 0, 3},
		},
		{
			give: give{
				nums: []int{7, 7, 7, 7},
			},
			want: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		got := smallerNumbersThanCurrent(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
