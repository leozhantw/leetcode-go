package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	s []byte
}

func Test_reverseString(t *testing.T) {
	tests := []struct {
		give give
		want []byte
	}{
		{
			give: give{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			give: give{
				s: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	for _, tt := range tests {
		got := reverseString(tt.give.s)
		assert.Equal(t, tt.want, got)
	}
}
