/*
 * @lc app=leetcode.cn id=383 lang=golang
 * @lcpr version=
 *
 * [383] 赎金信
 */
package main

// @lc code=start
// 思路：使用哈希表记录 magazine 中每个字符的出现次数
// 遍历 ransomNote 中的每个字符，检查是否能在 magazine 中找到足够的字符
// 时间复杂度：O(m+n)，空间复杂度：O(k)，其中 m 和 n 分别是 ransomNote 和 magazine 的长度，k 是字符集大小

func canConstruct(ransomNote string, magazine string) bool {
	// charCount 记录 magazine 中每个字符的出现次数
	charCount := map[rune]int{}

	// 统计 magazine 中每个字符的出现次数
	for _, char := range magazine {
		charCount[char]++
	}

	// 检查 ransomNote 中的每个字符是否能在 magazine 中找到
	for _, char := range ransomNote {
		if charCount[char] <= 0 {
			return false
		}
		charCount[char]--
	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "a"\n"b"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"ab"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"aab"\n
// @lcpr case=end

*/
