package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	root *TreeNode
}

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		give give
		want []int
	}{
		{
			give: give{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: []int{1, 3, 2},
		},
		{
			give: give{
				root: nil,
			},
			want: nil,
		},
		{
			give: give{
				root: &TreeNode{Val: 1},
			},
			want: []int{1},
		},
		{
			give: give{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
				},
			},
			want: []int{2, 1},
		},
		{
			give: give{
				root: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		got := inorderTraversal(tt.give.root)
		assert.Equal(t, tt.want, got)
	}
}
