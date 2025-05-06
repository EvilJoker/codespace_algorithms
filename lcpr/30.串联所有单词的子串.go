/*
 * @lc app=leetcode.cn id=30 lang=golang
 * @lcpr version=
 *
 * [30] 串联所有单词的子串
 */
package main

// @lc code=start
func findSubstring(s string, words []string) []int {
	// 解题思路：
	// 1. 计算所有单词的总长度，作为滑动窗口的大小
	// 2. 使用滑动窗口遍历字符串s，每次移动一个字符
	// 3. 对于每个窗口，检查是否包含所有单词（顺序可以不同）
	// 4. 使用map记录每个单词出现的次数，确保单词数量匹配
	// 时间复杂度：O(n*m*k)，其中n是s的长度，m是words的长度，k是单词的平均长度
	// 空间复杂度：O(m)，用于存储words的map

	sn := len(s)
	if sn == 0 || len(words) == 0 {
		return []int{}
	}

	// 计算所有单词的总长度
	sum := 0
	for _, w := range words {
		sum += len(w)
	}

	res := []int{}
	wordLen := len(words[0])

	// 滑动窗口遍历
	for i := 0; i <= sn-sum; i++ {
		if isSame(s[i:i+sum], words, wordLen) {
			res = append(res, i)
		}
	}

	return res
}

func isSame(s string, words []string, wordLen int) bool {
	// 使用map记录每个单词应该出现的次数
	dict := make(map[string]int)
	for _, w := range words {
		dict[w]++
	}

	// 检查每个单词是否出现正确的次数
	for i := 0; i < len(s); i += wordLen {
		cur := s[i : i+wordLen]
		dict[cur]--
		if dict[cur] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "barfoothefoobarman"\n["foo","bar"]\n
// @lcpr case=end

// @lcpr case=start
// "wordgoodgoodgoodbestword"\n["word","good","best","word"]\n
// @lcpr case=end

// @lcpr case=start
// "barfoofoobarthefoobarman"\n["bar","foo","the"]\n
// @lcpr case=end

*/
