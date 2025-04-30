/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// 动态规划，
	/*
		1. dp[i][j] word1[0:i] ---> word2[0:j]需要的变化的个数
		2. 初始化，当 i=0, 那么dp[0][j] = j
		3. 递推，当w1[i-1]==w2[j-1],dp[i][j]=dp[i-1][j-1]
		4. 不相等时，dp[i][j]= min(dp[i][j-1],dp[i-1][j]) +1, 三个情况
	*/
	m, n := len(word1), len(word2)

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 0
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]

			} else {

				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m][n]

}

// @lc code=end

/*
// @lcpr case=start
// "horse"\n"ros"\n
// @lcpr case=end

// @lcpr case=start
// "intention"\n"execution"\n
// @lcpr case=end

*/

