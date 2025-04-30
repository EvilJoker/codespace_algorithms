/*
 * @lc app=leetcode.cn id=236 lang=golang
 * @lcpr version=
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// sl: 包含 p. q 的最小树
	// 肯定不包含
	if root == nil {
		return nil
	}

	// 包含一个时
	if root == p || root == q {
		return root
	}
	// 如果包含连个
	left_node := lowestCommonAncestor(root.Left, p, q)
	right_node := lowestCommonAncestor(root.Right, p, q)

	// 肯定不包含的场景, 传递nil
	if left_node == nil && right_node == nil {
		return nil
	}
	// 只包含单节点, 传递目标节点
	if (left_node == p || left_node == q) && right_node == nil {
		return left_node
	}

	if (right_node == p || right_node == q) && left_node == nil {
		return right_node
	}

	// 两个都包含, 找到了，传递不是nil ，也不是目标接的节点
	if (left_node == p && right_node == q) || (right_node == p && left_node == q) {
		return root
	}

	// 直接向上传递
	if left_node != nil && left_node != p && left_node != q {
		return left_node
	}

	if right_node != nil && right_node != p && right_node != q {
		return right_node
	}

	return nil

}

// @lc code=end

/*
// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n1\n
// @lcpr case=end

// @lcpr case=start
// [3,5,1,6,2,0,8,null,null,7,4]\n5\n4\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n2\n
// @lcpr case=end

*/
