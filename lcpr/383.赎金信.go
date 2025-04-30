/*
 * @lc app=leetcode.cn id=383 lang=golang
 * @lcpr version=
 *
 * [383] 赎金信
 */
package main

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	//sl: 字段，判断 ransomNote 在不在 magazine 中
	dict := map[rune]int{}
	for _, c := range magazine {
		dict[c]++
	}

	for _, c := range ransomNote {
		if dict[c] <= 0 {
			return false
		}
		dict[c]--
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "a"\n"b"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"ab"\n
// @lcpr case=end

// @lcpr case=start
// "aa"\n"aab"\n
// @lcpr case=end

*/
