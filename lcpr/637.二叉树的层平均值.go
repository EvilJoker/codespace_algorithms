/*
 * @lc app=leetcode.cn id=637 lang=golang
 * @lcpr version=
 *
 * [637] 二叉树的层平均值
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
2. 使用两个队列：
   - nodeQueue: 存储节点
   - levelQueue: 存储对应的层级
3. 在遍历过程中：
   - 当遇到新的层级时，计算上一层的平均值
   - 累加当前层级的节点值和计数
4. 最后处理最后一层的平均值
*/
func averageOfLevels(root *TreeNode) []float64 {
	// 存储每层的平均值
	averages := []float64{}
	if root == nil {
		return averages
	}

	// 初始化队列
	nodeQueue := []*TreeNode{root}
	levelQueue := []int{0}
	currentLevel := 0
	levelSum, levelCount := 0, 0

	for len(nodeQueue) > 0 {
		// 取出队首节点和其层级
		node := nodeQueue[0]
		level := levelQueue[0]
		nodeQueue = nodeQueue[1:]
		levelQueue = levelQueue[1:]

		// 如果进入新的一层，计算上一层的平均值
		if level != currentLevel {
			if levelCount > 0 {
				averages = append(averages, float64(levelSum)/float64(levelCount))
			}
			levelSum, levelCount = 0, 0
			currentLevel = level
		}

		// 累加当前层的节点值和计数
		levelSum += node.Val
		levelCount++

		// 将子节点加入队列
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			levelQueue = append(levelQueue, level+1)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			levelQueue = append(levelQueue, level+1)
		}
	}

	// 处理最后一层的平均值
	if levelCount > 0 {
		averages = append(averages, float64(levelSum)/float64(levelCount))
	}

	return averages
}

// @lc code=end

/*
// @lcpr case=start
// [3,9,20,null,null,15,7]\n
// @lcpr case=end

// @lcpr case=start
// [3,9,20,15,7]\n
// @lcpr case=end

*/
