/*
 * @lc app=leetcode.cn id=104 lang=golang
 * @lcpr version=
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	return maxDep(root, 0)
}

func maxDep(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	return max(maxDep(root.Left, depth+1), maxDep(root.Right, depth+1))

}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,2]\n
// @lcpr case=end

*/
