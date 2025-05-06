/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=
 *
 * [64] 最小路径和
 */

package main

// @lc code=start
func minPathSum(grid [][]int) int {
	/* 解题思路：动态规划
	1. 定义dp数组：dp[i][j]表示从左上角(0,0)到位置(i,j)的最小路径和
	2. 初始化：
	   - dp[0][0] = grid[0][0]
	   - 第一行：dp[0][j] = dp[0][j-1] + grid[0][j]
	   - 第一列：dp[i][0] = dp[i-1][0] + grid[i][0]
	3. 状态转移：
	   - dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	   表示当前位置的最小路径和等于上方和左方较小值加上当前位置的值
	4. 最终结果：返回dp[m-1][n-1]
	*/

	// 获取网格的行数和列数
	m, n := len(grid), len(grid[0])

	// 初始化动态规划数组，dp[i][j]表示从(0,0)到(i,j)的最小路径和
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化起点
	dp[0][0] = grid[0][0]

	// 初始化第一列：只能从上往下走
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 初始化第一行：只能从左往右走
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// 动态规划递推：每个位置的最小路径和等于上方和左方较小值加上当前位置的值
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	// 返回右下角的最小路径和
	return dp[m-1][n-1]

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,1],[1,5,1],[4,2,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3],[4,5,6]]\n
// @lcpr case=end

*/
