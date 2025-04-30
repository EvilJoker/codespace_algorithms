/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=
 *
 * [392] 判断子序列
 */
package main

// @lc code=start
func isSubsequence(s string, t string) bool {

	sn, tn, p, q := len(s), len(t), 0, 0

	if s == "" {
		return true
	}

	for q < tn {
		if t[q] == s[p] {
			p++
			if p == sn {
				return true
			}
		}
		q++
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/
