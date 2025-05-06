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
/*
思路：
1. 使用层序遍历（BFS）遍历二叉树
2. 对于每一层，记录最后一个访问的节点，该节点即为右视图节点
3. 使用两个队列分别存储节点和对应的深度
4. 当深度变化时，将上一层的最后一个节点加入结果集
*/

func rightSideView(root *TreeNode) []int {
	// 存储右视图结果
	result := []int{}

	// 处理空树情况
	if root == nil {
		return result
	}

	// 初始化队列，存储节点和对应的深度
	nodeQueue := []*TreeNode{root}
	depthQueue := []int{0}

	// 记录上一层的深度和最后一个节点
	prevDepth := 0
	lastNode := root

	// 层序遍历
	for len(nodeQueue) > 0 {
		// 取出队首节点和深度
		currNode := nodeQueue[0]
		currDepth := depthQueue[0]

		// 更新队列
		nodeQueue = nodeQueue[1:]
		depthQueue = depthQueue[1:]

		// 当深度变化时，将上一层的最后一个节点加入结果集
		if prevDepth != currDepth {
			result = append(result, lastNode.Val)
			prevDepth = currDepth
		}
		lastNode = currNode

		// 将左右子节点加入队列
		if currNode.Left != nil {
			nodeQueue = append(nodeQueue, currNode.Left)
			depthQueue = append(depthQueue, currDepth+1)
		}
		if currNode.Right != nil {
			nodeQueue = append(nodeQueue, currNode.Right)
			depthQueue = append(depthQueue, currDepth+1)
		}
	}

	// 处理最后一层的最后一个节点
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
