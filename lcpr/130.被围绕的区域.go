/*
 * @lc app=leetcode.cn id=130 lang=golang
 * @lcpr version=
 *
 * [130] 被围绕的区域
 */
package main

// @lc code=start
/*
思路：
1. 从边界开始DFS，找到所有与边界相连的'O'，这些'O'不会被包围
2. 将这些不会被包围的'O'临时标记为'W'
3. 最后遍历整个board：
   - 将剩余的'O'（被包围的）改为'X'
   - 将'W'（不被包围的）改回'O'
*/

func solve(board [][]byte) {
	// 处理空board的情况
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])

	// DFS函数：标记与边界相连的'O'
	var markUnsurrounded func(int, int)
	markUnsurrounded = func(row, col int) {
		// 越界或不是'O'则返回
		if row < 0 || row >= rows || col < 0 || col >= cols || board[row][col] != 'O' {
			return
		}
		// 标记为不会被包围的'O'
		board[row][col] = 'W'
		// 向四个方向继续搜索
		markUnsurrounded(row-1, col) // 上
		markUnsurrounded(row+1, col) // 下
		markUnsurrounded(row, col-1) // 左
		markUnsurrounded(row, col+1) // 右
	}

	// 从边界开始DFS
	for i := 0; i < rows; i++ {
		markUnsurrounded(i, 0)      // 左边界
		markUnsurrounded(i, cols-1) // 右边界
	}
	for j := 0; j < cols; j++ {
		markUnsurrounded(0, j)      // 上边界
		markUnsurrounded(rows-1, j) // 下边界
	}

	// 处理所有格子
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X' // 被包围的'O'改为'X'
			case 'W':
				board[i][j] = 'O' // 不被包围的'W'改回'O'
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
