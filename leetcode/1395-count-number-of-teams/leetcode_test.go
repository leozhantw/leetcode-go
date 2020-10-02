package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	rating []int
}

func Test_numTeams(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				rating: []int{2, 5, 3, 4, 1},
			},
			want: 3,
		},
		{
			give: give{
				rating: []int{2, 1, 3},
			},
			want: 0,
		},
		{
			give: give{
				rating: []int{},
			},
			want: 0,
		},
		{
			give: give{
				rating: []int{1, 2, 3, 4},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		got := numTeams(tt.give.rating)
		assert.Equal(t, tt.want, got)
	}
}
