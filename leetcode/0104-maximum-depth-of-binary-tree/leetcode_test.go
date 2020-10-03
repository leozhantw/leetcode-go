package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	root *TreeNode
}

func Test_maxDepth(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				root: &TreeNode{},
			},
			want: 1,
		},
		{
			give: give{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: 3,
		},
		{
			give: give{
				root: &TreeNode{
					Val: 21,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		got := maxDepth(tt.give.root)
		assert.Equal(t, tt.want, got)
	}
}
