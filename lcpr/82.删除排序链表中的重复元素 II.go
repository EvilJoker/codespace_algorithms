/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=
 *
 * [82] 删除排序链表中的重复元素 II
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
func deleteDuplicates(head *ListNode) *ListNode {
	// sl： 完全去重,比如连续出现222， 那么就一个也不保留

	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	for cur != nil {
		// 出现重复
		if cur.Next != nil && cur.Val == cur.Next.Val {
			// 跳过所有重复的节点
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next // 连接到下一个不重复节点
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/
