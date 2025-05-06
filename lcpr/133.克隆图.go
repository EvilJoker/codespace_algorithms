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

/*
思路：
1. 使用深度优先搜索(DFS)遍历原图
2. 使用哈希表记录已克隆的节点，避免重复克隆
3. 对于每个节点：
   - 如果节点已被克隆，直接返回克隆节点
   - 如果节点未被克隆，创建新节点并递归克隆其邻居
4. 最终返回克隆后的图
*/

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// 记录已克隆的节点，key为原节点，value为克隆节点
	clonedNodes := map[*Node]*Node{}

	var dfs func(*Node) *Node
	dfs = func(curr *Node) *Node {
		// 如果节点已被克隆，直接返回
		if cloned, exists := clonedNodes[curr]; exists {
			return cloned
		}

		// 创建新节点
		newNode := &Node{Val: curr.Val}
		clonedNodes[curr] = newNode

		// 递归克隆所有邻居节点
		for _, neighbor := range curr.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, dfs(neighbor))
		}

		return newNode
	}

	return dfs(node)
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
