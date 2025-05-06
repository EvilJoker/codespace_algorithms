/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=
 *
 * [14] 最长公共前缀
 */
package main

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 解题思路：
	// 1. 从第一个字符串开始，逐个字符遍历，直到不符合条件
	// 2. 时间复杂度：O(n)，空间复杂度：O(1)

	res := ""
	index := 0
	for index < len(strs[0]) {
		for _, s := range strs {
			if index > len(s)-1 {
				return res
			}
			if s[index] != strs[0][index] {
				return res
			}
		}
		res += string(strs[0][index])
		index++

	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/
