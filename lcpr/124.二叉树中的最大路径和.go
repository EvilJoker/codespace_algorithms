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
	// 使用闭包来维护全局最大值
	maxSum := math.MinInt32

	// 定义递归函数，返回以当前节点为起点的最大路径和
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归获取左右子树的最大路径和，如果为负则取0
		leftSum := max(0, dfs(node.Left))
		rightSum := max(0, dfs(node.Right))

		// 更新全局最大值：当前节点作为路径中心的情况
		maxSum = max(maxSum, node.Val+leftSum+rightSum)

		// 返回以当前节点为起点的最大路径和（只能选择一边）
		return node.Val + max(leftSum, rightSum)
	}

	dfs(root)
	return maxSum
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
