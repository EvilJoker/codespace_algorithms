/*
 * @lc app=leetcode.cn id=199 lang=golang
 * @lcpr version=
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	//sl: 层序遍历，找到最右侧节点
	/*
		1. 使用队列
		2. 记录每个节点以及其层号，当自己是该层最后一个时，保存进结果

	*/

	result := []int{}

	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	queueDep := []int{0}
	lastdep := 0
	lastNode := root

	for len(queue) != 0 {
		head := queue[0]
		headDep := queueDep[0]

		queue = queue[1:]
		queueDep = queueDep[1:]

		if lastdep != headDep {
			lastdep = headDep
			result = append(result, lastNode.Val)
		}
		lastNode = head

		if head.Left != nil {
			queue = append(queue, head.Left)
			queueDep = append(queueDep, headDep+1)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
			queueDep = append(queueDep, headDep+1)
		}
	}
	result = append(result, lastNode.Val)
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,null,5,null,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,null,null,null,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,null,3]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
