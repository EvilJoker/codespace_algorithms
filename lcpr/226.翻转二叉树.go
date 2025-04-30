/*
 * @lc app=leetcode.cn id=226 lang=golang
 * @lcpr version=
 *
 * [226] 翻转二叉树
 */
package main

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left, right := root.Left, root.Right

	invertTree(left)
	invertTree(right)

	root.Left, root.Right = right, left

	return root

}

// @lc code=end

/*
// @lcpr case=start
// [4,2,7,1,3,6,9]\n
// @lcpr case=end

// @lcpr case=start
// [2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
