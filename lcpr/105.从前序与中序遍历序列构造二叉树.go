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
	// 前序遍历: 根 -> 左 -> 右
	// 中序遍历: 左 -> 根 -> 右
	// 1. 前序遍历的第一个元素是根节点
	// 2. 在中序遍历中找到根节点的位置，左边是左子树，右边是右子树
	// 3. 递归构建左右子树

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 获取根节点值
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 在中序遍历中找到根节点的位置
	rootIndex := findIndex(inorder, rootVal)

	// 递归构建左右子树
	// 左子树: 前序遍历[1:rootIndex+1], 中序遍历[0:rootIndex]
	// 右子树: 前序遍历[rootIndex+1:], 中序遍历[rootIndex+1:]
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func findIndex(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
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
