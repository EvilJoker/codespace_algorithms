/*
 * @lc app=leetcode.cn id=62 lang=golang
 * @lcpr version=30104
 *
 * [62] 不同路径
 */
package main

// @lc code=start
func uniquePaths(m int, n int) int {
	/* 解题思路：动态规划
	1. 问题分析：
	   - 机器人每次只能向下或向右移动一步
	   - 需要计算从左上角(0,0)到右下角(m-1,n-1)的所有可能路径数
	   - 这是一个典型的动态规划问题，因为每个位置的路径数依赖于其上方和左方的路径数

	2. 定义dp数组：
	   - dp[i][j]表示从起点(0,0)到位置(i,j)的不同路径数
	   - 数组维度为m×n，对应网格大小

	3. 初始化：
	   - 第一行和第一列的所有位置都只有1种路径（只能向右或向下）
	   - dp[0][j] = 1 (0 <= j < n)：第一行只能向右走
	   - dp[i][0] = 1 (0 <= i < m)：第一列只能向下走

	4. 状态转移方程：
	   - dp[i][j] = dp[i-1][j] + dp[i][j-1]
	   - 表示当前位置的路径数等于上方和左方路径数之和
	   - 因为到达当前位置只能从上方或左方移动一步

	5. 最终结果：
	   - 返回dp[m-1][n-1]，即到达右下角的路径数

	时间复杂度：O(m*n)，需要遍历整个网格
	空间复杂度：O(m*n)，需要存储dp数组
	*/

	// 创建并初始化dp数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化第一行和第一列
	// 第一行：只能向右移动
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	// 第一列：只能向下移动
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 动态规划递推：计算每个位置的路径数
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 当前位置的路径数 = 上方路径数 + 左方路径数
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 返回右下角的路径数
	return dp[m-1][n-1]
}

// @lc code=end

/*
// @lcpr case=start
// 3\n7\n
// @lcpr case=end

// @lcpr case=start
// 3\n2\n
// @lcpr case=end

// @lcpr case=start
// 7\n3\n
// @lcpr case=end

// @lcpr case=start
// 3\n3\n
// @lcpr case=end

*/
