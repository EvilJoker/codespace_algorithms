/*
 * @lc app=leetcode.cn id=151 lang=golang
 * @lcpr version=
 *
 * [151] 反转字符串中的单词
 */
package main

import "strings"

// @lc code=start
// 思路：
// 1. 使用 strings.Fields 分割字符串，自动处理多余空格
// 2. 从后向前遍历单词数组
// 3. 将单词拼接成新字符串，注意处理空格
// 4. 最后去除首尾空格

func reverseWords(s string) string {
	// 使用 Fields 分割字符串，自动处理多余空格
	words := strings.Fields(s)

	// 存储结果
	result := ""

	// 从后向前遍历单词数组
	for i := len(words) - 1; i >= 0; i-- {
		// 拼接单词和空格
		result += words[i] + " "
	}

	// 去除首尾空格并返回
	return strings.TrimSpace(result)
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
