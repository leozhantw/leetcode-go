package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	n     int
	start int
}

func Test_xorOperation(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				n:     5,
				start: 0,
			},
			want: 8,
		},
		{
			give: give{
				n:     4,
				start: 3,
			},
			want: 8,
		},
		{
			give: give{
				n:     1,
				start: 7,
			},
			want: 7,
		},
		{
			give: give{
				n:     10,
				start: 5,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		got := xorOperation(tt.give.n, tt.give.start)
		assert.Equal(t, tt.want, got)
	}
}
