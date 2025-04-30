package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// sl: 设置进位，节点依次相加。
	up := 0
	dummy := &ListNode{}
	pre := dummy
	for l1 != nil && l2 != nil {
		sum := up + l1.Val + l2.Val
		up = sum / 10
		val := sum % 10
		// 新节点
		pre.Next = &ListNode{Val: val}
		pre = pre.Next
		// 跟心
		l1 = l1.Next
		l2 = l2.Next
	}
	// 还有未结束的值

	for l1 != nil {
		sum := up + l1.Val
		up = sum / 10
		val := sum % 10
		pre.Next = &ListNode{Val: val}
		pre = pre.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sum := up + l2.Val
		up = sum / 10
		val := sum % 10
		pre.Next = &ListNode{Val: val}
		pre = pre.Next
		l2 = l2.Next
	}

	if up != 0 {
		pre.Next = &ListNode{Val: up}
	}

	return dummy.Next

}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/
