/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=
 *
 * [76] 最小覆盖子串
 */
package main

import "math"

// @lc code=start
func minWindow(s string, t string) string {

	/*
		解题思路：滑动窗口
		1. 使用map记录目标字符串t中每个字符的出现次数（初始为负值，表示"欠账"）
		2. 使用双指针维护一个滑动窗口：
		   - 右指针end向右移动，将遇到的字符加入window
		   - 当window中所有字符的计数都>=0时，说明当前窗口包含了t的所有字符
		   - 此时尝试收缩左指针start，直到不满足条件
		3. 在收缩过程中记录最小覆盖子串的起始位置和长度
		4. 时间复杂度：O(n)，空间复杂度：O(k)，其中n为s的长度，k为字符集大小
	*/
	// 初始化字符计数窗口
	window := make(map[rune]int)

	// 记录目标字符串中每个字符的"欠账"数量
	for _, char := range t {
		window[char]--
	}

	// 将输入字符串转换为rune切片以便处理Unicode字符
	sRunes := []rune(s)
	strLen := len(sRunes)

	// 初始化滑动窗口的指针和结果变量
	left, right := 0, 0
	minStart, minLen := 0, math.MaxInt32

	// 移动右指针扩展窗口
	for right < strLen {
		// 将右指针指向的字符加入窗口
		window[sRunes[right]]++

		// 如果当前窗口不满足条件，继续扩展
		if !isWindowValid(window) {
			right++
			continue
		}

		// 当窗口满足条件时，尝试收缩左边界
		for left <= right && isWindowValid(window) {
			// 更新最小覆盖子串
			windowSize := right - left + 1
			if windowSize < minLen {
				minStart = left
				minLen = windowSize
			}

			// 收缩左边界
			window[sRunes[left]]--
			left++
		}
		right++
	}

	// 处理未找到有效解的情况
	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}

// isWindowValid 检查当前窗口是否包含所有目标字符
func isWindowValid(window map[rune]int) bool {
	for _, count := range window {
		if count < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
