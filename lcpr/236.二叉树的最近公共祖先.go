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
/*
思路：
1. 递归遍历二叉树，对于每个节点，判断其是否包含p或q
2. 如果当前节点为nil，返回nil
3. 如果当前节点是p或q，返回当前节点
4. 递归处理左右子树
5. 根据左右子树的返回值判断：
   - 如果左右子树都返回nil，说明当前子树不包含p和q
   - 如果左右子树只有一个返回非nil，说明找到了p或q中的一个
   - 如果左右子树都返回非nil，说明找到了p和q，当前节点就是最近公共祖先
*/

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 基本情况：空节点或找到目标节点
	if root == nil || root == p || root == q {
		return root
	}

	// 递归处理左右子树
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左右子树都找到了目标节点，当前节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 如果只有左子树找到目标节点，返回左子树的结果
	if left != nil {
		return left
	}

	// 如果只有右子树找到目标节点，返回右子树的结果
	return right
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
