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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}

	p1, p2 := list1, list2

	pre := dummy
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
