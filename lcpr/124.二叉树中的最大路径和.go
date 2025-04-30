/*
 * @lc app=leetcode.cn id=124 lang=golang
 * @lcpr version=
 *
 * [124] 二叉树中的最大路径和
 */
package main

import (
	"math"
)

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	var maxsum = math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		// 当前节点能作为一条完整路径的“中心”
		current := root.Val + max(0, left) + max(0, right)
		maxsum = max(maxsum, current)

		// 向父节点返回的贡献值（只能选择一边）
		return root.Val + max(0, max(left, right))
	}
	dfs(root)
	return maxsum
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [-10,9,20,null,null,15,7]\n
// @lcpr case=end

*/
