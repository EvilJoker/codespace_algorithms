/*
 * @lc app=leetcode.cn id=205 lang=golang
 * @lcpr version=
 *
 * [205] 同构字符串
 */
package main

// @lc code=start
// 思路：使用两个哈希表记录字符的映射关系
// 1. 如果两个字符串长度不同，直接返回 false
// 2. 遍历字符串，用哈希表记录每个字符最后出现的位置
// 3. 如果当前字符在各自哈希表中的位置不同，说明映射关系不一致，返回 false
// 4. 使用 i+1 作为位置值，避免与 map 默认值 0 冲突
func isIsomorphic(s string, t string) bool {
	// s2t 记录 s 到 t 的映射位置
	// t2s 记录 t 到 s 的映射位置
	s2t, t2s := map[byte]int{}, map[byte]int{}

	// 长度不同，直接返回 false
	if len(s) != len(t) {
		return false
	}

	// 遍历字符串，检查映射关系
	for i := 0; i < len(s); i++ {
		sChar, tChar := s[i], t[i]
		// 如果映射位置不一致，说明不是同构字符串
		if s2t[sChar] != t2s[tChar] {
			return false
		}
		// 更新字符最后出现的位置（使用 i+1 避免与默认值 0 冲突）
		s2t[sChar], t2s[tChar] = i+1, i+1
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
