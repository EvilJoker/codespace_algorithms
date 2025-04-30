/*
 * @lc app=leetcode.cn id=114 lang=golang
 * @lcpr version=
 *
 * [114] 二叉树展开为链表
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

func flatten(root *TreeNode) {
	//sl: 前序遍历 root,保存在数组，然后拼接
	nodes := &[]*TreeNode{}
	preorder(root, nodes)
	dummy := &TreeNode{Right: root}
	for _, node := range *nodes {
		dummy.Right = node
		dummy.Left = nil
		dummy = dummy.Right
	}
}

func preorder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root)
	// 先序遍历
	preorder(root.Left, nodes)
	preorder(root.Right, nodes)
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,5,3,4,null,6]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [0]\n
// @lcpr case=end

*/
