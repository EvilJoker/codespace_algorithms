/*
 * @lc app=leetcode.cn id=108 lang=golang
 * @lcpr version=
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	//sl: 二叉搜索树，左 < 根 < 右
	// 初始化一个节点，然后不断新增节点，维护进去 -- 非平衡二叉是
	// 平衡二叉树， 使用低估，每次将中间节点作为新的根节点

	var dfs func(start, end int) *TreeNode

	dfs = func(start, end int) *TreeNode {
		// 结束匹配
		if start > end {
			return nil
		}
		mid := (start-end)/2 + end // 中间节点作为最高节点

		root := &TreeNode{Val: nums[mid]}
		root.Left = dfs(start, mid-1)
		root.Right = dfs(mid+1, end)

		return root
	}

	return dfs(0, len(nums)-1)

}

// @lc code=end

/*
// @lcpr case=start
// [-10,-3,0,5,9]\n
// @lcpr case=end

// @lcpr case=start
// [1,3]\n
// @lcpr case=end

*/
