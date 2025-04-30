/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */
package main

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	//sl: 层序遍历

	if root == nil {
		return nil
	}

	queue := []*Node{root}
	depthQueue := []int{0}

	lastNode := root
	lastDepth := 0

	for len(queue) != 0 {
		head, depth := queue[0], depthQueue[0]
		queue = queue[1:] // 弹出
		depthQueue = depthQueue[1:]

		// 拼接
		if depth != lastDepth {
			lastNode.Next = nil
		} else {
			lastNode.Next = head
		}
		lastNode, lastDepth = head, depth

		// 放入
		if head.Left != nil {
			queue = append(queue, head.Left)
			depthQueue = append(depthQueue, depth+1)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
			depthQueue = append(depthQueue, depth+1)
		}

	}
	lastNode.Next = nil // 最后一层末尾节点置空
	return root

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
