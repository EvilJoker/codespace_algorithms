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
func averageOfLevels(root *TreeNode) []float64 {
	// sl:还是层序遍历
	result := []float64{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	queueDep := []int{0}
	lastDep := 0
	count, sum := 0, 0

	for len(queue) > 0 {
		head, dep := queue[0], queueDep[0]
		queue = queue[1:]
		queueDep = queueDep[1:]
		if dep != lastDep {

			if count == 0 {
				result = append(result, 0)
			} else {
				result = append(result, float64(sum)/float64(count))
			}
			count, sum = 0, 0
		}
		lastDep = dep
		count++
		sum += head.Val

		if head.Left != nil {
			queue = append(queue, head.Left)
			queueDep = append(queueDep, dep+1)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
			queueDep = append(queueDep, dep+1)
		}
	}
	if count == 0 {
		result = append(result, 0)
	} else {
		result = append(result, float64(sum)/float64(count))
	}
	return result

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
