/*
 * @lc app=leetcode.cn id=72 lang=golang
 * @lcpr version=
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	/* 解题思路：动态规划
	1. 定义dp数组：dp[i][j]表示将word1的前i个字符转换为word2的前j个字符所需的最少操作数
	2. 初始化：
	   - dp[0][0] = 0 (空字符串转换为空字符串不需要操作)
	   - dp[i][0] = i (将word1的前i个字符转换为空字符串需要i次删除操作)
	   - dp[0][j] = j (将空字符串转换为word2的前j个字符需要j次插入操作)
	3. 状态转移：
	   - 当word1[i-1] == word2[j-1]时：dp[i][j] = dp[i-1][j-1] (不需要操作)
	   - 当word1[i-1] != word2[j-1]时：
	     - 插入：dp[i][j-1] + 1 (理解：dp[i][j-1]表示已经将word1的前i个字符转换为word2的前j-1个字符，然后插入word2[j-1])
	     - 删除：dp[i-1][j] + 1 (理解：将word1的前i-1个字符转换为word2的前j个字符，然后删除word1[i-1])
	     - 替换：dp[i-1][j-1] + 1 (理解：将word1的前i-1个字符转换为word2的前j-1个字符，然后将word1[i-1]替换为word2[j-1])
	     - 取三者最小值：dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
	4. 最终结果：返回dp[m][n]，其中m和n分别为word1和word2的长度
	*/
	// 获取两个字符串的长度
	m, n := len(word1), len(word2)

	// 初始化动态规划数组，dp[i][j]表示将word1前i个字符转换为word2前j个字符所需的最少操作数
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// 初始化基本情况
	dp[0][0] = 0 // 空字符串转换为空字符串不需要操作

	// 初始化第一列：将word1前i个字符转换为空字符串需要i次删除操作
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	// 初始化第一行：将空字符串转换为word2前j个字符需要j次插入操作
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	// 动态规划递推
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				// 当前字符相等，不需要操作，直接继承前一个状态
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 当前字符不相等，取三种操作中的最小值：
				// 1. 删除：dp[i-1][j] + 1
				// 2. 插入：dp[i][j-1] + 1
				// 3. 替换：dp[i-1][j-1] + 1
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	// 返回将word1转换为word2所需的最少操作数
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

