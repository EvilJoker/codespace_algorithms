/*
 * @lc app=leetcode.cn id=290 lang=golang
 * @lcpr version=
 *
 * [290] 单词规律
 */
package main

import "strings"

// @lc code=start
// 思路：使用双向映射判断模式匹配
// 1. 将字符串s按空格分割成单词数组
// 2. 使用两个map分别记录pattern字符和单词的映射关系
// 3. 遍历pattern和单词数组，确保每个字符和单词的映射关系一致
// 4. 使用索引+1作为映射值，避免0值冲突
func wordPattern(pattern string, s string) bool {
	// pattern字符到索引的映射
	charToIndex := map[byte]int{}
	// 单词到索引的映射
	wordToIndex := map[string]int{}

	// 将字符串s分割成单词数组
	words := strings.Fields(s)

	// 检查pattern和单词数组长度是否一致
	if len(pattern) != len(words) {
		return false
	}

	// 遍历pattern和单词数组
	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		// 如果当前字符和单词的映射值不一致，说明不匹配
		if charToIndex[char] != wordToIndex[word] {
			return false
		}

		// 更新映射值（使用索引+1避免0值冲突）
		charToIndex[char] = i + 1
		wordToIndex[word] = i + 1
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dog cat cat dog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dog cat cat fish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dog cat cat dog"\n
// @lcpr case=end

*/
