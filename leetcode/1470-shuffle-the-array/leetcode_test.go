package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
	n    int
}

func Test_shuffle(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			give: give{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			give: give{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			want: []int{1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		got := shuffle(tt.give.nums, tt.give.n)
		assert.Equal(t, tt.want, got)
	}
}
