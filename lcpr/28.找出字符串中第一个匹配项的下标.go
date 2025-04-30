/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=
 *
 * [28] 找出字符串中第一个匹配项的下标
 */
package main

import "strings"

// @lc code=start
func strStr(haystack string, needle string) int {
	// 直接用库函数
	return strings.Index(haystack, needle)

}

// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/
