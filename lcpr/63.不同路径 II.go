/*
 * @lc app=leetcode.cn id=63 lang=golang
 * @lcpr version=
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*sl: 动态规划
	1. dp[i][j] 表示到达当前位置的路径数
	2. 初始值，obstacleGrid[i][j] ==1, 时 dp[i][j]=0,其他时间，都是 左边 +1
	3 递推obstacleGrid[i][j] ==0,dp[i][j]= dp[i-1][j]+ dp[i][j-1] 线路数是加的关系
	*/
	row, col := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	// 初始化
	if obstacleGrid[0][0] == 0 {

		dp[0][0] = 1
	}
	// 中间有障碍物就过不去了
	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 0 && dp[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}
	for j := 1; j < col; j++ {
		if obstacleGrid[0][j] == 0 && dp[0][j-1] == 1 {
			dp[0][j] = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			//递归
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

