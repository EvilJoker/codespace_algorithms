/*
 * @lc app=leetcode.cn id=221 lang=golang
 * @lcpr version=
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {

	/*
		解题思路：动态规划
		1. 定义dp数组：dp[i][j]表示以(i,j)为右下角的最大正方形的边长
		2. 初始化：
		   - 第一行和第一列：如果matrix[i][j]='1'，则dp[i][j]=1
		   - 其他位置初始化为0
		3. 状态转移：
		   - 当matrix[i][j]='1'时：
		     dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
		     表示当前位置的最大正方形边长取决于其上方、左方和左上方的三个位置的最小值加1
		4. 最终结果：返回最大边长的平方
	*/

	// 获取矩阵的行数和列数
	row := len(matrix)
	col := len(matrix[0])

	// 初始化动态规划数组，dp[i][j]表示以(i,j)为右下角的最大正方形的边长
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	// 记录最大正方形的边长
	mx_side := 0

	// 动态规划递推
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 处理第一行和第一列的情况
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				}
			} else if matrix[i][j] == '0' {
				// 当前位置为0，不能构成正方形
				dp[i][j] = 0
			} else {
				// 当前位置为1，取左、上、左上三个位置的最小值加1
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}

			// 更新最大边长
			if dp[i][j] != 0 && mx_side < dp[i][j] {
				mx_side = dp[i][j]
			}
		}
	}

	// 返回最大正方形的面积
	return mx_side * mx_side

}

// @lc code=end

/*
// @lcpr case=start
// [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0","1"],["1","0"]]\n
// @lcpr case=end

// @lcpr case=start
// [["0"]]\n
// @lcpr case=end

*/

