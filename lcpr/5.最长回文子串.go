package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	/* 解题思路：动态规划
	1. 定义dp数组：dp[i][j]表示字符串s从位置i到位置j的子串是否为回文串
	2. 初始化：
	   - 单个字符一定是回文串：dp[i][i] = true
	   - 相邻字符：dp[i][i+1] = s[i] == s[i+1]
	3. 状态转移：
	   - 当s[i] == s[j]时：
	     - 如果j-i=1，说明是相邻字符，直接判断是否相等
	     - 否则，dp[i][j] = dp[i+1][j-1]
	4. 遍历顺序：
	   - 外层循环遍历子串长度
	   - 内层循环遍历起始位置
	5. 记录结果：
	   - 在遍历过程中记录最长的回文子串的起始和结束位置
	*/

	// 获取字符串长度
	n := len(s)

	// 初始化动态规划数组，dp[i][j]表示s[i..j]是否为回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	// 记录最长回文子串的起始位置和长度
	maxlen := 0
	left := 0
	right := 0

	// 动态规划填充dp数组
	// 外层循环j表示子串的结束位置
	// 内层循环i表示子串的起始位置，且i <= j
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				// 单个字符一定是回文串
				dp[i][j] = true
			} else if j == i+1 {
				// 相邻字符，直接判断是否相等
				dp[i][j] = s[i] == s[j]
			} else {
				// 状态转移：当前字符相等且内部子串是回文
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			// 如果当前子串是回文且长度大于已知最大长度，更新结果
			if dp[i][j] && maxlen < j-i+1 {
				left = i
				right = j
				maxlen = j - i + 1
			}
		}
	}

	// 返回最长回文子串
	return s[left : right+1]
}

// @lc code=end

/*
// @lcpr case=start
// "babad"\n
// @lcpr case=end

// @lcpr case=start
// "cbbd"\n
// @lcpr case=end

*/
