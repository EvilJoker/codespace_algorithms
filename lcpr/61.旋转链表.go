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
	//sl: 类似旋转数组，但是旋转链表更简单，找打头尾节点进行拼接。 将末尾k 个节点迁移
	if head == nil || k == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	k = k % length

	if head == nil || length == 0 || k == 0 {
		return head // ✅ 直接返回原链表
	}

	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
		}
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	newtail := slow
	newhead := slow.Next
	oldtail := fast

	oldtail.Next = head
	newtail.Next = nil

	return newhead
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
