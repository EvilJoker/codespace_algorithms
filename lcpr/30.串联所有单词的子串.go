/*
 * @lc app=leetcode.cn id=30 lang=golang
 * @lcpr version=
 *
 * [30] 串联所有单词的子串
 */
package main

// @lc code=start
func findSubstring(s string, words []string) []int {
	// sl: 计算 words 中总长，然后 从 s 开始遍历。判断 words 是不是 s 的子串
	sn := len(s)

	sum := 0
	for _, w := range words {
		sum += len(w)

	}
	res := []int{}

	for i := 0; i <= sn-sum; i++ {
		if isSame(s[i:i+sum], words, len(words[0])) {
			res = append(res, i)
		}
	}

	return res
}

func isSame(s string, words []string, wordlen int) bool {

	dict := make(map[string]int)

	for _, w := range words {
		dict[w]++
	}

	for i := 0; i < len(s); i += wordlen {
		cur := s[i : i+wordlen]

		dict[cur]--
		if dict[cur] < 0 {
			return false
		}

	}

	return true
}

// @lc code=end

/*
// @lcpr case=start
// "barfoothefoobarman"\n["foo","bar"]\n
// @lcpr case=end

// @lcpr case=start
// "wordgoodgoodgoodbestword"\n["word","good","best","word"]\n
// @lcpr case=end

// @lcpr case=start
// "barfoofoobarthefoobarman"\n["bar","foo","the"]\n
// @lcpr case=end

*/
