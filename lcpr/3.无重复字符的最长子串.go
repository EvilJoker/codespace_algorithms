// package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=
 *
 * [3] 无重复字符的最长子串
 */

package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	/*
		解题思路：滑动窗口
		1. 使用双指针维护一个滑动窗口，窗口内字符不重复
		2. 使用哈希表记录窗口内字符是否存在
		3. 右指针不断向右移动，当遇到重复字符时，左指针向右移动直到不重复
		4. 每次更新窗口大小时，记录最大长度
		时间复杂度：O(n)，空间复杂度：O(字符集大小)
	*/

	// 将字符串转换为rune数组，以支持Unicode字符
	sr := []rune(s)
	n := len(sr)

	// 初始化变量
	start := 0                    // 窗口左边界
	maxLen := 0                   // 最大长度
	window := make(map[rune]bool) // 记录窗口内字符

	// 右指针不断向右移动
	for end := 0; end < n; end++ {
		char := sr[end]

		// 当遇到重复字符时，移动左指针直到不重复
		for window[char] {
			// 移除左指针指向的字符
			window[sr[start]] = false
			start++
		}

		// 将当前字符加入窗口
		window[char] = true

		// 更新最大长度
		currLen := end - start + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}

// @lc code=end

/*
// @lcpr case=start
// "abcabcbb"\n
// @lcpr case=end

// @lcpr case=start
// "bbbbb"\n
// @lcpr case=end

// @lcpr case=start
// "pwwkew"\n
// @lcpr case=end

*/
