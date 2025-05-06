/*
 * @lc app=leetcode.cn id=530 lang=golang
 * @lcpr version=
 *
 * [530] 二叉搜索树的最小绝对差
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "math"

// 思路：
// 1. 二叉搜索树的中序遍历结果是一个升序数组
// 2. 最小绝对差一定出现在相邻节点之间
// 3. 使用中序遍历，记录前一个节点，计算当前节点与前一个节点的差值
// 4. 维护最小差值

func getMinimumDifference(root *TreeNode) int {
	// 初始化最小差值为最大整数
	minDiff := math.MaxInt16
	// 记录前一个访问的节点
	var prevNode *TreeNode

	// 定义中序遍历函数
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 先遍历左子树
		inorder(node.Left)

		// 计算当前节点与前一个节点的差值
		if prevNode != nil {
			minDiff = min(node.Val-prevNode.Val, minDiff)
		}

		// 更新前一个节点
		prevNode = node

		// 遍历右子树
		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,6,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,48,null,null,12,49]\n
// @lcpr case=end

*/
