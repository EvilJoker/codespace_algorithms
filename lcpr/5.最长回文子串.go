package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 * @lcpr version=
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	/* 思路：
	1. 暴力：从头到尾，以每个元素为中心点。判断是否回文字符串。
	2. 优化：动态规划。d[i][j] 标识 i~j 是否是回文字符串
	3. 初始化，i=j,dp[i][j] =true
	4. 递归：si==sj, 且 dp[i+1][j-1]=true，那么dp[i][j ]也是真的
	 或者 j=i+1
	 // 边界： 遍历时 考虑相邻情况，j 在

	*/
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	maxlen := 0
	left := 0
	right := 0

	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if i == j {
				// 赋值
				dp[i][j] = true
			} else if j == i+1 {
				dp[i][j] = s[i] == s[j]

			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}

			// 为true 的都判断下是否更新
			if dp[i][j] && maxlen < j-i+1 {
				left = i
				right = j
				maxlen = j - i + 1
			}

		}
	}

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
