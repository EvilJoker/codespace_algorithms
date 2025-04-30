/*
 * @lc app=leetcode.cn id=49 lang=golang
 * @lcpr version=
 *
 * [49] 字母异位词分组
 */
package main

import "sort"

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	//sl: 将排序后一样的字符串作为一组

	dict := map[string]([]string){}

	for _, s := range strs {
		runes := []rune(s)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		key := string(runes)

		dict[key] = append(dict[key], s)
	}

	res := [][]string{}
	for key := range dict {
		res = append(res, dict[key])
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["eat", "tea", "tan", "ate", "nat", "bat"]\n
// @lcpr case=end

// @lcpr case=start
// [""]\n
// @lcpr case=end

// @lcpr case=start
// ["a"]\n
// @lcpr case=end

*/
