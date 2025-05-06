/*
 * @lc app=leetcode.cn id=139 lang=golang
 * @lcpr version=30200
 *
 * [139] 单词拆分
 *
 * https://leetcode.cn/problems/word-break/description/
 *
 * algorithms
 * Medium (58.10%)
 * Likes:    2747
 * Dislikes: 0
 * Total Accepted:    784K
 * Total Submissions: 1.3M
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
 *
 * 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
 *
 *
 *
 * 示例 1：
 *
 * 输入: s = "leetcode", wordDict = ["leet", "code"]
 * 输出: true
 * 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
 *
 *
 * 示例 2：
 *
 * 输入: s = "applepenapple", wordDict = ["apple", "pen"]
 * 输出: true
 * 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
 * 注意，你可以重复使用字典中的单词。
 *
 *
 * 示例 3：
 *
 * 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * 输出: false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 300
 * 1 <= wordDict.length <= 1000
 * 1 <= wordDict[i].length <= 20
 * s 和 wordDict[i] 仅由小写英文字母组成
 * wordDict 中的所有字符串 互不相同
 *
 *
 */
package main

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	// 动态规划解法思路：
	// 1. 定义dp数组：dp[i]表示字符串s的前i个字符是否可以被wordDict中的单词拼接而成
	// 2. 初始化：dp[0] = true，表示空字符串可以被拼接
	// 3. 状态转移：
	//    - 遍历字符串s的每个位置i
	//    - 对于每个位置i，检查所有可能的单词长度j
	//    - 如果dp[i-j]为true且s[i-j:i]在wordDict中，则dp[i] = true
	// 4. 最终结果：返回dp[len(s)]
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			word_len := len(word)
			if i >= word_len && dp[i-word_len] {
				if s[i-word_len:i] == word {
					dp[i] = true
					break
				}
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
