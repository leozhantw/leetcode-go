package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type give struct {
	n int
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		give give
		want []string
	}{
		{
			give: give{
				n: 15,
			},
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
		{
			give: give{
				n: 1,
			},
			want: []string{"1"},
		},
	}

	for _, tt := range tests {
		got := fizzBuzz(tt.give.n)
		assert.Equal(t, tt.want, got)
	}
}
