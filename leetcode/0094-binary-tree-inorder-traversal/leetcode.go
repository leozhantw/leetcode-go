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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}

func traverse(node *TreeNode, res *[]int) {
	if node != nil {
		traverse(node.Left, res)
		*res = append(*res, node.Val)
		traverse(node.Right, res)
	}
}
