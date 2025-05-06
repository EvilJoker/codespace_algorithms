/*
 * @lc app=leetcode.cn id=97 lang=golang
 * @lcpr version=
 *
 * [97] 交错字符串
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	/*
		解题思路：动态规划
		1. 定义dp数组：dp[i][j]表示s1的前i个字符和s2的前j个字符是否能交错组成s3的前i+j个字符
		2. 初始化：
		   - dp[0][0] = true (空字符串可以组成空字符串)
		   - 第一行：dp[i][0] = s1[0:i] == s3[0:i]
		   - 第一列：dp[0][j] = s2[0:j] == s3[0:j]
		3. 状态转移：
		   - 如果s1的第i个字符等于s3的第i+j个字符，则dp[i][j] = dp[i-1][j]
		   - 如果s2的第j个字符等于s3的第i+j个字符，则dp[i][j] = dp[i][j-1]
		   - 两种情况满足其一即可
	*/
	// 获取三个字符串的长度
	m, n, p := len(s1), len(s2), len(s3)
	// 如果s1和s2的长度之和不等于s3的长度，直接返回false
	if m+n != p {
		return false
	}

	// 初始化动态规划数组，dp[i][j]表示s1前i个字符和s2前j个字符是否能交错组成s3前i+j个字符
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// 初始化基本情况
	dp[0][0] = true // 空字符串可以组成空字符串

	// 初始化第一列：只用s1的前i个字符匹配s3的前i个字符
	for i := 1; i <= m; i++ {
		dp[i][0] = s1[i-1] == s3[i-1] && dp[i-1][0]
	}

	// 初始化第一行：只用s2的前j个字符匹配s3的前j个字符
	for j := 1; j <= n; j++ {
		dp[0][j] = s2[j-1] == s3[j-1] && dp[0][j-1]
	}

	// 动态规划递推
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果s1的第i个字符等于s3的第i+j个字符，且s1前i-1个字符和s2前j个字符能组成s3前i+j-1个字符
			if s1[i-1] == s3[i+j-1] && dp[i-1][j] {
				dp[i][j] = true
				continue
			}

			// 如果s2的第j个字符等于s3的第i+j个字符，且s1前i个字符和s2前j-1个字符能组成s3前i+j-1个字符
			if s2[j-1] == s3[i+j-1] && dp[i][j-1] {
				dp[i][j] = true
			}
		}
	}

	// 返回最终结果：s1和s2是否能交错组成s3
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

