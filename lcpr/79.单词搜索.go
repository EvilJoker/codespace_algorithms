/*
 * @lc app=leetcode.cn id=79 lang=golang
 * @lcpr version=
 *
 * [79] 单词搜索
 */
package main

// @lc code=start
func exist(board [][]byte, word string) bool {
	/*
		解题思路：回溯算法
		1. 问题分析：
		   - 在二维网格中搜索单词
		   - 单词必须由相邻单元格的字母按顺序组成
		   - 同一个单元格内的字母不允许被重复使用
		   - 可以从任意位置开始搜索

		2. 算法设计：
		   - 使用DFS回溯算法进行搜索
		   - 对每个单元格作为起点进行尝试
		   - 搜索方向：上下左右四个方向
		   - 使用访问标记避免重复使用同一单元格

		3. 复杂度分析：
		   - 时间复杂度：O(m*n*4^L)，其中m*n为网格大小，L为单词长度
		   - 空间复杂度：O(L)，递归调用栈的深度

		4. 优化点：
		   - 使用访问标记避免重复访问
		   - 及时剪枝：不匹配时立即返回
	*/
	// 获取网格的行数、列数和目标单词的长度
	row, col, n := len(board), len(board[0]), len(word)

	// 定义四个搜索方向：下、上、右、左
	directs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	// 定义DFS函数，用于在网格中搜索单词
	// i, j: 当前搜索位置的坐标
	// index: 当前匹配到单词的第几个字符
	var dfs func(i, j int, index int) bool

	dfs = func(i, j int, index int) bool {
		// 如果已经匹配完所有字符，说明找到了单词
		if index == n {
			return true
		}

		// 检查是否超出网格边界
		if i < 0 || i > row-1 || j < 0 || j > col-1 {
			return false
		}

		// 检查当前字符是否匹配且未被访问过
		if board[i][j] != word[index] {
			return false
		}

		// 标记当前位置为已访问
		board[i][j] = '@'

		// 尝试四个方向的搜索
		for _, direct := range directs {
			x, y := i+direct[0], j+direct[1]
			if dfs(x, y, index+1) {
				return true
			}
		}

		// 回溯：恢复当前位置的原始字符
		board[i][j] = word[index]

		return false
	}

	// 遍历网格中的每个位置作为起始点
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
