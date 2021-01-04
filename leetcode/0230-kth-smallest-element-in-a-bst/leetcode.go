package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var count, res int
	inorder(root, k, &count, &res)
	return res
}

func inorder(node *TreeNode, k int, count, res *int) {
	if node == nil {
		return
	}

	inorder(node.Left, k, count, res)
	*count++
	if *count == k {
		*res = node.Val
		return
	}
	inorder(node.Right, k, count, res)
}
