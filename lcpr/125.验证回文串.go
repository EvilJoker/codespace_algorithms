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
/*
思路：
1. 将字符串转换为小写，统一大小写处理
2. 使用双指针从两端向中间移动
3. 跳过非字母数字字符
4. 比较左右指针指向的字符是否相同
5. 如果不同则不是回文串，如果全部相同则是回文串
*/
func isPalindrome(s string) bool {
	// 转换为小写，统一大小写处理
	s = strings.ToLower(s)

	// 转换为rune切片，以正确处理Unicode字符
	chars := []rune(s)
	left, right := 0, len(chars)-1

	// 双指针从两端向中间移动
	for left < right {
		// 跳过左侧非字母数字字符
		if !unicode.IsDigit(chars[left]) && !unicode.IsLetter(chars[left]) {
			left++
			continue
		}
		// 跳过右侧非字母数字字符
		if !unicode.IsDigit(chars[right]) && !unicode.IsLetter(chars[right]) {
			right--
			continue
		}
		// 比较左右指针指向的字符
		if chars[left] != chars[right] {
			return false
		}
		// 移动指针
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
