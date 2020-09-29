package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_numIdenticalPairs(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				nums: []int{1, 2, 3, 1, 1, 3},
			},
			want: 4,
		},
		{
			give: give{
				nums: []int{1, 1, 1, 1},
			},
			want: 6,
		},
		{
			give: give{
				nums: []int{1, 2, 3},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := numIdenticalPairs(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
