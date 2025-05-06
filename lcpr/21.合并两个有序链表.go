/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=
 *
 * [21] 合并两个有序链表
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

// 解题思路：
// 1. 使用虚拟头节点 dummy 来简化链表操作
// 2. 使用双指针 p1, p2 分别指向两个链表的当前节点
// 3. 比较两个指针指向节点的值，将较小的节点连接到结果链表中
// 4. 移动指针和结果链表的指针
// 5. 最后处理剩余节点
// 时间复杂度：O(n+m)，空间复杂度：O(1)

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建虚拟头节点
	dummy := &ListNode{}

	// 初始化两个指针
	p1, p2 := list1, list2

	// pre 指向当前结果链表的末尾
	pre := dummy

	// 当两个链表都不为空时，比较节点值并合并
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			pre.Next = p1
			p1 = p1.Next
		} else {
			pre.Next = p2
			p2 = p2.Next
		}
		pre = pre.Next
	}

	// 处理剩余节点
	if p1 != nil {
		pre.Next = p1
	}
	if p2 != nil {
		pre.Next = p2
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
