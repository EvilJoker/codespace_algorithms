/*
 * @lc app=leetcode.cn id=141 lang=golang
 * @lcpr version=
 *
 * [141] 环形链表
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

/*
思路：
1. 使用快慢指针法检测链表是否有环
2. 快指针每次移动两步，慢指针每次移动一步
3. 如果链表有环，快慢指针最终会相遇
4. 如果快指针到达链表末尾，说明链表无环
*/
func hasCycle(head *ListNode) bool {
	// 初始化快慢指针都指向头节点
	fast, slow := head, head

	// 快指针每次移动两步，需要确保fast和fast.Next都不为空
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next // 快指针移动两步
		slow = slow.Next      // 慢指针移动一步

		// 如果快慢指针相遇，说明链表有环
		if fast == slow {
			return true
		}
	}

	// 快指针到达链表末尾，说明链表无环
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

*/
