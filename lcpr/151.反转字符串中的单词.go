/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=
 *
 * [151] 反转字符串中的单词
 */
package main

import "strings"

// @lc code=start
func reverseWords(s string) string {
	// sl: 先 spilit ,然后 翻转
	strs := strings.Fields(s)

	res := ""

	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != " " {
			res += strings.TrimSpace(strs[i]) + " "
		}
	}
	return strings.TrimSpace(res)

}

// @lc code=end

/*
// @lcpr case=start
// "the sky is blue"\n
// @lcpr case=end

// @lcpr case=start
// "  hello world  "\n
// @lcpr case=end

// @lcpr case=start
// "a good   example"\n
// @lcpr case=end

*/
