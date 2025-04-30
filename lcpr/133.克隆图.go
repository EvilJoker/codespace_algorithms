/*
 * @lc app=leetcode.cn id=133 lang=golang
 * @lcpr version=
 *
 * [133] 克隆图
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
package main

func cloneGraph(node *Node) *Node {
	// sl: 使用dfs ，第一次克隆所有节点，存入 map。完成遍历后。进行第二次遍历，将原先的 neughbors,节点使用 map 替换；。
	// 改进： 不需要遍历两次, 使用 dfs, 当遍历到没有被clone 的节点，递归遍历。直到所有节点都是 克隆后的节点
	if node == nil {
		return nil
	}
	visted := map[*Node]*Node{}

	var cloneNode func(head *Node) *Node

	cloneNode = func(head *Node) *Node {
		// clone 节点
		newNode := &Node{Val: head.Val}
		visted[head] = newNode

		cloneNeighbors := []*Node{}

		for _, neighbor := range head.Neighbors {

			cloned, exist := visted[neighbor]
			if exist {
				cloneNeighbors = append(cloneNeighbors, cloned)
			} else {
				clone_neighbor := cloneNode(neighbor)
				cloneNeighbors = append(cloneNeighbors, clone_neighbor)
			}

		}
		// 复制所有邻接节点
		newNode.Neighbors = cloneNeighbors

		return newNode

	}

	return cloneNode(node)

}

// @lc code=end

/*
// @lcpr case=start
// [[2,4],[1,3],[2,4],[1,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
