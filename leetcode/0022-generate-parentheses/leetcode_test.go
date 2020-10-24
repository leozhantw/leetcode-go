package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	n int
}

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		give give
		want []string
	}{
		{
			give: give{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			give: give{
				n: 1,
			},
			want: []string{"()"},
		},
	}

	for _, tt := range tests {
		got := generateParenthesis(tt.give.n)
		assert.Equal(t, tt.want, got)
	}
}
