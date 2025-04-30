/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	//sl: 回溯，上下左右，每个都是候选项。当找不到就结束。每一个元素都是起点
	row, col, n := len(board), len(board[0]), len(word)

	directs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(i, j int, index int) bool

	dfs = func(i, j int, index int) bool {
		if index == n { // 已经找完了
			return true
		}

		if i < 0 || i > row-1 || j < 0 || j > col-1 {
			return false // 超出边界
		}

		// 不匹配或者访问过
		if board[i][j] != word[index] {
			return false
		}

		board[i][j] = '@' // 访问标记

		for _, direct := range directs {
			x, y := i+direct[0], j+direct[1]
			if dfs(x, y, index+1) {
				return true
			}
		}

		// 回溯

		board[i][j] = word[index]

		return false
	}

	// 每个节点都遍历
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false

}

// @lc code=end

/*
// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"SEE"\n
// @lcpr case=end

// @lcpr case=start
// [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCB"\n
// @lcpr case=end

*/
