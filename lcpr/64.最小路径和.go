/*
 * @lc app=leetcode.cn id=64 lang=golang
 * @lcpr version=
 *
 * [64] 最小路径和
 */

package main

// @lc code=start
func minPathSum(grid [][]int) int {
	// 思路：二维动态规划
	/*
		1. dp[i][j] 从左上角触发，到当前的最短路径
		2. 初始值： dp[i][0]，都是横向或竖向的值
		3. 递推： dp[i][j]= min(dp[i-1][j],dp[i][j-1])+ grid[i][j]
	*/

	row := len(grid)
	col := len(grid[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	// 初始值
	dp[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	//递推

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {

			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]

		}
	}

	return dp[row-1][col-1]

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
