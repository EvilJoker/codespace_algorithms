/*
 * @lc app=leetcode.cn id=125 lang=golang
 * @lcpr version=
 *
 * [125] 验证回文串
 */
package main

import (
	"strings"
	"unicode"
)

// @lc code=start
func isPalindrome(s string) bool {
	// *去除无效字符串

	s = strings.ToLower(s)

	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left < right {
		if !unicode.IsDigit(runes[left]) && !unicode.IsLetter(runes[left]) {
			left++
			continue
		}
		if !unicode.IsDigit(runes[right]) && !unicode.IsLetter(runes[right]) {
			right--
			continue
		}
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true

}

// @lc code=end

/*
// @lcpr case=start
// "A man, a plan, a canal: Panama"\n
// @lcpr case=end

// @lcpr case=start
// "race a car"\n
// @lcpr case=end

// @lcpr case=start
// " "\n
// @lcpr case=end

*/
