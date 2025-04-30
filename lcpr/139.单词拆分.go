/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	/* 思路：动规，
	1. dp[i]bool, 0~i字符 是否是
	2. 初始值dp[0]=0
	3. 递推。range dict 中的单词，看是否 dp[i-len(w)]是否是
	*/
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, w := range wordDict {
			// 单词相等，上一个也相等
			if i-len(w) >= 0 && dp[i-len(w)] && s[i-len(w):i] == w {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]

}

// @lc code=end

/*
// @lcpr case=start
// "leetcode"\n["leet", "code"]\n
// @lcpr case=end

// @lcpr case=start
// "applepenapple"\n["apple", "pen"]\n
// @lcpr case=end

// @lcpr case=start
// "catsandog"\n["cats", "dog", "sand", "and", "cat"]\n
// @lcpr case=end

*/

