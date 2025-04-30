/*
 * @lc app=leetcode.cn id=221 lang=golang
 * @lcpr version=
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {

	/*
		sl: 最大正方形
		1.dp[i][j] = 变长。以ij 为右下角的最大正方形
		2.初始值，if m[i][j]=1,那么 max 就是 1
		3. 递推 m[i][j]=1,dp=max(三个) +1
	*/

	row := len(matrix)
	col := len(matrix[0])

	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}

	mx_side := 0

	// 递推
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				}

			} else if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}

			if dp[i][j] != 0 && mx_side < dp[i][j] {
				mx_side = dp[i][j]
			}
		}
	}

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

