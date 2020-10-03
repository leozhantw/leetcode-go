package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			give: give{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		got := singleNumber(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
