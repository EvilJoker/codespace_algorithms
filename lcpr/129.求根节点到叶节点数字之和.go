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
func sumNumbers(root *TreeNode) int {
	// sl: 计算根到每一个叶子节点

	return sumn(root, 0)

}

func sumn(root *TreeNode, num int) int {
	if root == nil {
		// 如果叶子节点没有，这一次的相加就不存在
		return 0
	}

	cur := num*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return cur
	}

	return sumn(root.Left, cur) + sumn(root.Right, cur)

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
