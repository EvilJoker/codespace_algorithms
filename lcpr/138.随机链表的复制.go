/*
 * @lc app=leetcode.cn id=138 lang=golang
 * @lcpr version=
 *
 * [138] 随机链表的复制
 */
package main

type node struct {
	Val    int
	Next   *node
	Random *node
}

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *node) *node {
	// sl: 使用hash 两次遍历
	dummy := &node{Next: head}

	dict := map[*node]*node{}

	pre := dummy
	cur := head

	for cur != nil {
		newnode := &node{Val: cur.Val, Next: cur.Next, Random: cur.Random}
		dict[cur] = newnode
		cur = cur.Next
		pre.Next = newnode
		pre = pre.Next
	}

	pre = dummy.Next
	for pre != nil {
		if pre.Random != nil {
			pre.Random = dict[pre.Random]
		}
		pre = pre.Next
	}

	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [[7,null],[13,0],[11,4],[10,2],[1,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,1],[2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[3,null],[3,0],[3,null]]\n
// @lcpr case=end

*/
