/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=
 *
 * [92] 反转链表 II
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	/*
		解题思路：
		1. 使用虚拟头节点简化操作
		2. 找到需要反转区间的前一个节点
		3. 使用头插法进行区间反转
		4. 保持其他节点相对位置不变
	*/
	dummy := &ListNode{Next: head}
	pre := dummy
	// 寻找前一个节点
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	cur := pre.Next // 起点
	var next *ListNode
	// 记忆方法：这个是一个循环。 cur,nex, pre
	for i := 0; i < right-left; i++ {

		next = cur.Next      // 抓头
		cur.Next = next.Next // 摘叶

		next.Next = pre.Next //扔头

		pre.Next = next //插尾

	}

	return dummy.Next

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
