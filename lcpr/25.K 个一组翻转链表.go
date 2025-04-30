/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	/*sl:
	1. 预处理 ，k 求余数
	2. 不满 k 就不翻转
	3. 保留前头后尾进行翻转
	*/

	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	length := 0

	for cur != nil {
		length++
		cur = cur.Next
	}

	// 翻转

	for length >= k {
		cur = pre.Next
		var next *ListNode

		for i := 1; i < k; i++ {
			next = cur.Next
			cur.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
		}

		//

		pre = cur // 下一个的新头
		length -= k
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
