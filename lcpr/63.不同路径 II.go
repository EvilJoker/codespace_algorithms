/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=
 *
 * [63] 不同路径 II
 */
package main

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*
	   动态规划解法:
	   1. 状态定义: dp[i][j] 表示从起点(0,0)到达位置(i,j)的不同路径数量
	   2. 状态转移:
	      - 如果当前位置是障碍物(obstacleGrid[i][j] == 1),则 dp[i][j] = 0
	      - 否则 dp[i][j] = dp[i-1][j] + dp[i][j-1] (从上方和左方到达的路径数之和)
	   3. 初始状态:
	      - dp[0][0] = 1 (如果起点不是障碍物)
	      - 第一行和第一列需要特殊处理,如果遇到障碍物则后续位置都不可达
	*/
	row, col := len(obstacleGrid), len(obstacleGrid[0])

	// 初始化dp数组
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}

	// 处理起点
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}

	// 处理第一列
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}

	// 处理第一行
	for j := 1; j < col; j++ {
		if obstacleGrid[0][j] == 0 && dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	// 填充dp数组
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[row-1][col-1]
}

// @lc code=end

/*
// @lcpr case=start
// [[0,0,0],[0,1,0],[0,0,0]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1],[0,0]]\n
// @lcpr case=end

*/
