/*
 * @lc app=leetcode.cn id=97 lang=golang
 * @lcpr version=
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	/*
		1. 动态规划：注意点，交错是指两个字符串，顺序不能变
		2. dp[][] bool, 表示 s1,前i 和 s2 前j, 组成了 s3 的 i+j 个. i 是下标，得 +1
		3. 初始化，如果 j=0, means: s1[0:i+1]==s3[0:i+1]
		4. 递推：dp[i][j]= dp[i][j-1] && s3[i-1+j] == s2[j-1] ||  dp[i-1][j] && s3[i-1+j] == s1[i-1]
	*/
	m, n, p := len(s1), len(s2), len(s3)
	if m+n != p {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	// 初始化
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = s1[i-1] == s3[i-1] && dp[i-1][0]
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = s2[j-1] == s3[j-1] && dp[0][j-1]
	}
	// 递推
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
			}

			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[m][n]

}

// @lc code=end

/*
// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbcbcac"\n
// @lcpr case=end

// @lcpr case=start
// "aabcc"\n"dbbca"\n"aadbbbaccc"\n
// @lcpr case=end

// @lcpr case=start
// ""\n""\n""\n
// @lcpr case=end

*/

