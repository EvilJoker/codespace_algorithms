/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=
 *
 * [242] 有效的字母异位词
 */
package main

// @lc code=start
// 思路：使用哈希表记录字符频率
// 1. 如果两个字符串长度不同，直接返回false
// 2. 遍历第一个字符串，统计每个字符出现的次数
// 3. 遍历第二个字符串，每遇到一个字符就减少对应计数
// 4. 如果某个字符的计数小于0，说明第二个字符串中该字符出现次数过多，返回false
// 5. 最后所有字符计数都为0，说明是有效的字母异位词

func isAnagram(s string, t string) bool {
	// 使用哈希表记录字符频率
	charCount := map[rune]int{}

	// 长度不同直接返回false
	if len(s) != len(t) {
		return false
	}

	// 统计第一个字符串中每个字符的出现次数
	for _, char := range s {
		charCount[char]++
	}

	// 遍历第二个字符串，减少对应字符的计数
	for _, char := range t {
		charCount[char]--
		// 如果某个字符计数小于0，说明不是有效的字母异位词
		if charCount[char] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/
