package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SubrectangleQueries(t *testing.T) {
	subrectangleQueries := Constructor([][]int{
		{1, 2, 1},
		{4, 3, 4},
		{3, 2, 1},
		{1, 1, 1},
	})

	t.Run("should get expected value", func(t *testing.T) {
		got := subrectangleQueries.GetValue(0, 2)
		assert.Equal(t, 1, got)
	})

	t.Run("update value in the range", func(t *testing.T) {
		subrectangleQueries.UpdateSubrectangle(0, 0, 2, 1, 5)
		assert.Equal(t, 5, subrectangleQueries.GetValue(0, 0))
		assert.Equal(t, 5, subrectangleQueries.GetValue(1, 1))
		assert.Equal(t, 5, subrectangleQueries.GetValue(2, 1))
		assert.Equal(t, 1, subrectangleQueries.GetValue(3, 0))
	})
}
