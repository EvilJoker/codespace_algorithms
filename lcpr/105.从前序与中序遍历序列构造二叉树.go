/*
 * @lc app=leetcode.cn id=105 lang=golang
 * @lcpr version=
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	//sl: 中序遍历确定分割点。 前序后续，确定头节点
	n := len(inorder)
	if n <= 0 {
		return nil
	}
	root := preorder[0]

	indexIn := findInex(inorder, root)

	leftLength := indexIn

	leftNode, rightNode := buildTree(preorder[1:leftLength+1], inorder[0:leftLength]),
		buildTree(preorder[leftLength+1:n], inorder[leftLength+1:n])

	rootNode := &TreeNode{Val: root, Left: leftNode, Right: rightNode}
	return rootNode

}

func findInex(order []int, target int) int {
	for i, v := range order {
		if v == target {
			return i
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,15,7]\n[9,3,15,20,7]\n
// @lcpr case=end

// @lcpr case=start
// [-1]\n[-1]\n
// @lcpr case=end

*/
