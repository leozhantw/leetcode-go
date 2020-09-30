package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums  []int
	index []int
}

func Test_createTargetArray(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums:  []int{0, 1, 2, 3, 4},
				index: []int{0, 1, 2, 2, 1},
			},
			want: []int{0, 4, 1, 3, 2},
		},
		{
			give: give{
				nums:  []int{1, 2, 3, 4, 0},
				index: []int{0, 1, 2, 3, 0},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			give: give{
				nums:  []int{1},
				index: []int{0},
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		got := createTargetArray(tt.give.nums, tt.give.index)
		assert.Equal(t, tt.want, got)
	}
}
