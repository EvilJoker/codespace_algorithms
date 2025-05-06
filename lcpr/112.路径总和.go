/*
 * @lc app=leetcode.cn id=112 lang=golang
 * @lcpr version=
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	/*
	   解题思路：
	   1. 使用递归遍历二叉树
	   2. 每次递归时，用目标值减去当前节点的值
	   3. 当到达叶子节点时，判断剩余值是否为0
	   4. 如果存在一条路径使得路径和等于目标值，返回true
	*/
	if root == nil {
		return false
	}

	// 到达叶子节点时判断
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	// 递归判断左右子树
	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}

// @lc code=end

/*
// @lcpr case=start
// [5,4,8,11,null,13,4,7,2,null,null,null,1]\n22\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3]\n5\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
