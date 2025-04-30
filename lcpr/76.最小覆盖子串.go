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

	// sl:滑动窗口. 先 根据t构造一个欠账的 map， 如果加入新元素，使得每项值都 >0 ，就说明包含
	window := map[rune]int{}

	// 构造一个负值 window
	for _, c := range t {
		window[c]--

	}

	sr := []rune(s)

	n := len(sr)
	start, end := 0, 0
	minstart, minend, minlen := 0, 0, math.MaxInt32

	for ; end < n; end++ {
		end_c := sr[end]
		// 填充新元素
		window[end_c]++

		// 不存在，继续右移
		if !contain(window) {
			continue
		}

		// 当 window 满足条件时，尝试收缩左边界
		for start <= end && contain(window) {
			if end-start+1 < minlen {
				// 更新逻辑
				minstart = start
				minlen = end - start + 1
			}
			start_c := sr[start]
			window[start_c]--
			start++
		}
	}
	if minlen == math.MaxInt32 {
		return ""
	}
	return s[minstart : minend+1]
}

// 不包含
func contain(dict map[rune]int) bool {
	for _, v := range dict {
		if v < 0 {
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
