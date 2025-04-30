/*
 * @lc app=leetcode.cn id=222 lang=golang
 * @lcpr version=
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	// sl: 普通思路是穷举遍历
	if root == nil {
		return 0
	}

	return countNodes(root.Left) + 1 + countNodes(root.Right)

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
