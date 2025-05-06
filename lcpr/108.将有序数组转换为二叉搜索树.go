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
	/*
	   解题思路：
	   1. 由于输入数组是有序的，要构建平衡二叉搜索树
	   2. 每次选择数组中间元素作为根节点，可以保证左右子树节点数平衡
	   3. 递归构建左右子树
	*/

	var buildBST func(left, right int) *TreeNode
	buildBST = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		// 选择中间元素作为根节点
		mid := left + (right-left)/2

		// 构建当前节点
		root := &TreeNode{Val: nums[mid]}

		// 递归构建左右子树
		root.Left = buildBST(left, mid-1)
		root.Right = buildBST(mid+1, right)

		return root
	}

	return buildBST(0, len(nums)-1)
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
