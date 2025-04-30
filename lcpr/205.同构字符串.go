/*
 * @lc app=leetcode.cn id=205 lang=golang
 * @lcpr version=
 *
 * [205] 同构字符串
 */
package main

// @lc code=start
func isIsomorphic(s string, t string) bool {
	//sl: 两个 map, 互相保存对方的 index, 映射的字符应该想能
	dict1, dict2 := map[byte]int{}, map[byte]int{}

	n1, n2 := len(s), len(t)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		c1, c2 := s[i], t[i]
		if dict1[c1] != dict2[c2] {
			return false
		}
		// 对于未出现过的字符，map 的默认值为 0，
		// 为避免 0 和合法索引 0 的冲突，我们用 i+1 来记录位置
		dict1[c1], dict2[c2] = i+1, i+1

	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// "egg"\n"add"\n
// @lcpr case=end

// @lcpr case=start
// "foo"\n"bar"\n
// @lcpr case=end

// @lcpr case=start
// "paper"\n"title"\n
// @lcpr case=end

*/
