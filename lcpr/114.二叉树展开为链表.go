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
	/*
	   解题思路：
	   1. 使用前序遍历收集所有节点
	   2. 将收集到的节点重新连接成链表
	   3. 每个节点的左子树置为nil，右子树指向下一个节点
	*/
	if root == nil {
		return
	}

	// 收集前序遍历的节点
	nodes := make([]*TreeNode, 0)
	preorder(root, &nodes)

	// 重新连接节点
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
	// 处理最后一个节点
	nodes[len(nodes)-1].Left = nil
	nodes[len(nodes)-1].Right = nil
}

func preorder(root *TreeNode, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	*nodes = append(*nodes, root)
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
