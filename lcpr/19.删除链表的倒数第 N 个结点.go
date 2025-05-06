/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 解题思路：
	// 1. 使用快慢指针，快指针先走n步，然后快慢指针同时移动
	// 2. 当快指针到达链表末尾时，慢指针指向倒数第n+1个节点
	// 3. 删除倒数第n个节点，即将慢指针的Next指向Next.Next
	// 4. 使用虚拟头节点处理删除头节点的情况
	// 5. 时间复杂度：O(n)，空间复杂度：O(1)

	// 创建虚拟头节点，处理删除头节点的情况
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy

	// 快指针先走n步
	for i := 0; i < n; i++ {
		if fast == nil {
			return dummy.Next
		}
		fast = fast.Next
	}

	// 快慢指针同时移动，直到快指针到达链表末尾
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 删除倒数第n个节点
	slow.Next = slow.Next.Next

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/
