/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=
 *
 * [130] 被围绕的区域
 */
package main

// @lc code=start
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'W' // 临时标记为不会被包围的
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	// 从边界的 O 开始 DFS
	for i := 0; i < rows; i++ {
		dfs(i, 0)
		dfs(i, cols-1)
	}
	for j := 0; j < cols; j++ {
		dfs(0, j)
		dfs(rows-1, j)
	}

	// 再遍历整个 board，处理剩下的
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X' // 被包围的
			} else if board[i][j] == 'W' {
				board[i][j] = 'O' // 不被包围的
			}
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]\n
// @lcpr case=end

// @lcpr case=start
// [["X"]]\n
// @lcpr case=end

*/
