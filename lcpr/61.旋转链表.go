/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	// 思路：
	// 1. 特殊情况处理：空链表或k=0时直接返回
	// 2. 计算链表长度，并对k取模（因为旋转长度超过链表长度时是循环的）
	// 3. 使用快慢指针找到需要断开的位置：
	//    - 快指针先走k步
	//    - 然后快慢指针同时移动，直到快指针到达末尾
	// 4. 重新连接链表：
	//    - 原末尾节点指向原头节点
	//    - 新末尾节点指向nil
	//    - 返回新的头节点

	// 特殊情况处理
	if head == nil || k == 0 {
		return head
	}

	// 计算链表长度
	length := 0
	for curr := head; curr != nil; curr = curr.Next {
		length++
	}

	// 对k取模，避免不必要的旋转
	k = k % length
	if k == 0 {
		return head
	}

	// 使用快慢指针找到断开位置
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// 快指针先走k步
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	// 快慢指针同时移动，直到快指针到达末尾
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 重新连接链表
	newHead := slow.Next // 新的头节点
	slow.Next = nil      // 新的尾节点指向nil
	fast.Next = head     // 原尾节点指向原头节点

	return newHead
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/
