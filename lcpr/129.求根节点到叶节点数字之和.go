/*
 * @lc app=leetcode.cn id=129 lang=golang
 * @lcpr version=
 *
 * [129] 求根节点到叶节点数字之和
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
// 思路：
// 1. 使用深度优先搜索(DFS)遍历二叉树
// 2. 从根节点开始，每遍历一个节点，就将当前路径上的数字累加
// 3. 当到达叶子节点时，将完整的路径数字加入结果
// 4. 使用递归实现，每次递归时传递当前路径的数字

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// dfs 函数用于深度优先搜索
// node: 当前节点
// currentSum: 当前路径的数字和
func dfs(node *TreeNode, currentSum int) int {
	// 空节点返回0
	if node == nil {
		return 0
	}

	// 计算当前路径的数字
	currentSum = currentSum*10 + node.Val

	// 如果是叶子节点，返回当前路径的数字
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// 递归计算左右子树的路径和
	return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,9,0,5,1]\n
// @lcpr case=end

*/
