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
	/*
		解题思路：
		1. 使用虚拟头节点简化链表操作
		2. 设置进位变量up，初始为0
		3. 同时遍历两个链表，对应节点相加并处理进位
		4. 处理较长的链表剩余部分
		5. 最后检查是否还有进位需要处理
	*/

	// 创建虚拟头节点，简化链表操作
	dummy := &ListNode{}
	pre := dummy
	up := 0 // 进位变量

	// 同时遍历两个链表，对应节点相加
	for l1 != nil && l2 != nil {
		sum := up + l1.Val + l2.Val
		up = sum / 10                  // 计算进位
		val := sum % 10                // 计算当前位值
		pre.Next = &ListNode{Val: val} // 创建新节点
		pre = pre.Next                 // 移动指针
		l1 = l1.Next                   // 移动链表指针
		l2 = l2.Next
	}

	// 处理l1剩余部分
	for l1 != nil {
		sum := up + l1.Val
		up = sum / 10
		val := sum % 10
		pre.Next = &ListNode{Val: val}
		pre = pre.Next
		l1 = l1.Next
	}

	// 处理l2剩余部分
	for l2 != nil {
		sum := up + l2.Val
		up = sum / 10
		val := sum % 10
		pre.Next = &ListNode{Val: val}
		pre = pre.Next
		l2 = l2.Next
	}

	// 处理最后的进位
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
