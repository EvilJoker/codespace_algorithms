/*
 * @lc app=leetcode.cn id=86 lang=golang
 * @lcpr version=
 *
 * [86] 分隔链表
 */
package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	/*
	   解题思路：
	   1. 将链表分成两部分：小于x的节点和大于等于x的节点
	   2. 使用两个虚拟头节点分别连接这两部分
	   3. 最后将两部分链表连接起来

	   优化点：
	   1. 使用虚拟头节点简化操作
	   2. 原地修改链表，不需要额外空间
	   3. 保持相对顺序
	*/

	// 创建两个虚拟头节点
	lessHead := &ListNode{}
	greaterHead := &ListNode{}
	less, greater := lessHead, greaterHead

	// 遍历原链表进行分割
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}

	// 连接两部分链表
	less.Next = greaterHead.Next
	greater.Next = nil

	return lessHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

*/
