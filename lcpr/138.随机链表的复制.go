/*
 * @lc app=leetcode.cn id=138 lang=golang
 * @lcpr version=
 *
 * [138] 随机链表的复制
 */
package main

type node struct {
	Val    int
	Next   *node
	Random *node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *node) *node {
	// 创建虚拟头节点，简化链表操作
	dummy := &node{Next: head}
	// 存储原节点到新节点的映射关系
	nodeMap := map[*node]*node{}

	// 第一次遍历：创建新节点
	prev := dummy
	curr := head
	for curr != nil {
		// 创建新节点，复制值
		newNode := &node{Val: curr.Val}
		// 建立映射关系
		nodeMap[curr] = newNode
		// 更新链表
		prev.Next = newNode
		prev = newNode
		curr = curr.Next
	}

	// 第二次遍历：设置Random指针
	curr = dummy.Next
	for curr != nil {
		if curr.Random != nil {
			// 通过映射关系找到对应的新节点
			curr.Random = nodeMap[curr.Random]
		}
		curr = curr.Next
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [[7,null],[13,0],[11,4],[10,2],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,null],[3,0],[3,null]]\n
// @lcpr case=end

*/
