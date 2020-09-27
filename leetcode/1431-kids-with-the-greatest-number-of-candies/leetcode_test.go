package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	candies      []int
	extraCandies int
}

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		give give
		want []bool
	}{
		{
			give: give{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			give: give{
				candies:      []int{4, 2, 1, 1, 2},
				extraCandies: 1,
			},
			want: []bool{true, false, false, false, false},
		},
		{
			give: give{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		got := kidsWithCandies(tt.give.candies, tt.give.extraCandies)
		assert.Equal(t, tt.want, got)
	}
}
