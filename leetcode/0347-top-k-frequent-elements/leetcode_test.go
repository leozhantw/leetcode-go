package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type give struct {
	nums []int
	k    int
}

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			give: give{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		got := topKFrequent(tt.give.nums, tt.give.k)
		assert.Equal(t, tt.want, got)
	}
}
