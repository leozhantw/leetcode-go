package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	nums []int
}

func Test_decompressRLElist(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{2, 4, 4, 4},
		},
		{
			give: give{
				nums: []int{1, 1, 2, 3},
			},
			want: []int{1, 3, 3},
		},
	}

	for _, tt := range tests {
		got := decompressRLElist(tt.give.nums)
		assert.Equal(t, tt.want, got)
	}
}
