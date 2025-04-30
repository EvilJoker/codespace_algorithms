/*
 * @lc app=leetcode.cn id=242 lang=golang
 * @lcpr version=
 *
 * [242] 有效的字母异位词
 */
package main

// @lc code=start
func isAnagram(s string, t string) bool {

	// sl: 判断字符频率是否一致
	dict := map[rune]int{}

	if len(s) != len(t) {
		return false
	}

	for _, c := range s {
		dict[c]++
	}

	for _, c := range t {
		dict[c]--

		if dict[c] < 0 {
			return false
		}

	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "anagram"\n"nagaram"\n
// @lcpr case=end

// @lcpr case=start
// "rat"\n"car"\n
// @lcpr case=end

*/
