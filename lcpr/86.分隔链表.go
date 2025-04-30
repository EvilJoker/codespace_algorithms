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
	//sl: 设计两个新头，遍历 head 进行分割，然后拼接
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	pre1, pre2 := dummy1, dummy2

	for head != nil {
		if head.Val < x {
			pre1.Next = head
			pre1 = pre1.Next
		} else {
			pre2.Next = head
			pre2 = pre2.Next
		}
		head = head.Next
	}
	pre1.Next = dummy2.Next
	pre2.Next = nil

	return dummy1.Next

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
