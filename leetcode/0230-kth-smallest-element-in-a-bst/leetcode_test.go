package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type give struct {
	root *TreeNode
	k    int
}

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		give give
		want int
	}{
		{
			give: give{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				k: 1,
			},
			want: 1,
		},
		{
			give: give{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
							Left: &TreeNode{
								Val: 1,
							},
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				k: 3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		got := kthSmallest(tt.give.root, tt.give.k)
		assert.Equal(t, tt.want, got)
	}
}
