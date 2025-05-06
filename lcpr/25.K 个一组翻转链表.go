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
	// 解题思路：
	// 1. 使用虚拟头节点简化操作
	// 2. 先计算链表长度，确定需要翻转的组数
	// 3. 对每组进行翻转，每组翻转k-1次
	// 4. 每组翻转后更新pre指针指向下一组的起始位置
	// 时间复杂度：O(n)，空间复杂度：O(1)

	// 创建虚拟头节点
	dummy := &ListNode{Next: head}
	pre := dummy
	cur := head

	// 计算链表长度
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	// 当剩余长度大于等于k时，进行翻转
	for length >= k {
		cur = pre.Next
		var next *ListNode

		// 每组需要翻转k-1次
		for i := 1; i < k; i++ {
			// 保存下一个节点
			next = cur.Next
			// 将当前节点指向下下个节点
			cur.Next = next.Next
			// 将下一个节点指向pre的下一个节点
			next.Next = pre.Next
			// 更新pre的下一个节点
			pre.Next = next
		}

		// 更新pre指针到下一组的起始位置
		pre = cur
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
