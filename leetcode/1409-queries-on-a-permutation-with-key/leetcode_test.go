package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	queries []int
	m       int
}

func Test_processQueries(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				queries: []int{3, 1, 2, 1},
				m:       5,
			},
			want: []int{2, 1, 2, 1},
		},
		{
			give: give{
				queries: []int{4, 1, 2, 2},
				m:       4,
			},
			want: []int{3, 1, 2, 0},
		},
		{
			give: give{
				queries: []int{7, 5, 5, 8, 3},
				m:       8,
			},
			want: []int{6, 5, 0, 7, 5},
		},
	}

	for _, tt := range tests {
		got := processQueries(tt.give.queries, tt.give.m)
		assert.Equal(t, tt.want, got)
	}
}
