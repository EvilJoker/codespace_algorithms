/*
 * @lc app=leetcode.cn id=290 lang=golang
 * @lcpr version=
 *
 * [290] 单词规律
 */
package main

import "strings"

// @lc code=start
func wordPattern(pattern string, s string) bool {
	// sl: 同构字符串，双向映射
	dict1, dict2 := map[byte]int{}, map[string]int{}

	strs := strings.Fields(s)

	n1, n2 := len(pattern), len(strs)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		c1, s1 := pattern[i], strs[i]
		if dict1[c1] != dict2[s1] {
			return false
		}
		dict1[c1], dict2[s1] = i+1, i+1
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "abba"\n"dog cat cat dog"\n
// @lcpr case=end

// @lcpr case=start
// "abba"\n"dog cat cat fish"\n
// @lcpr case=end

// @lcpr case=start
// "aaaa"\n"dog cat cat dog"\n
// @lcpr case=end

*/
