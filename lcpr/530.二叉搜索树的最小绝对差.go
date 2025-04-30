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

func getMinimumDifference(root *TreeNode) int {
	//sl : 2叉搜索树的概念： 父节点大于左、小于右 --z中序遍历是一个升序数组，比较当前节点和上一个节点差值即可

	minDif := math.MaxInt16
	var pre *TreeNode
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		if pre != nil {
			minDif = min(node.Val-pre.Val, minDif)

		}
		pre = node
		inorder(node.Right)

	}
	inorder(root)

	return minDif

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
