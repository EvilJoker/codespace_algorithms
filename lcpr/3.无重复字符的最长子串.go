// package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 * @lcpr version=
 *
 * [3] 无重复字符的最长子串
 */

package main

import "math"

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	/* 思路:滑动窗口,双指针不断滑动
	1. 维护 hash 判断字符串是否重复
	2. 当 右指针新加入元素，导致重复时。不断收缩左节点，直到不重复
	3. 维护 max 值
	*/
	sr := []rune(s)
	n := len(sr)
	start, end, max_len := 0, 0, math.MinInt32
	window := make(map[rune]bool)

	for ; end < n; end++ {
		// 加入新元素
		end_c := sr[end]

		for window[end_c] && start <= end { // 包含重复元素

			// 去除新元素
			start_c := sr[start]
			window[start_c] = false
			start++

		}

		max_len = max(max_len, end-start+1) // 更新
		window[end_c] = true

	}

	if max_len == math.MinInt32 {
		return 0
	}

	return max_len

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
