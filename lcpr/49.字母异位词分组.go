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
	// 思路：
	// 1. 字母异位词的特点是包含相同的字母，只是顺序不同
	// 2. 将每个字符串排序后作为key，原字符串作为value存入map
	// 3. 排序后的字符串相同的，说明是字母异位词，会被分到同一组
	// 4. 最后将map中的value转换为二维数组返回

	// 使用map存储分组结果，key为排序后的字符串，value为原字符串数组
	dict := make(map[string][]string)

	// 遍历每个字符串
	for _, s := range strs {
		// 将字符串转换为rune切片以便排序
		runes := []rune(s)
		// 对rune切片进行排序
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		// 将排序后的rune切片转回字符串作为key
		key := string(runes)
		// 将原字符串添加到对应的分组中
		dict[key] = append(dict[key], s)
	}

	// 将map中的分组结果转换为二维数组
	result := make([][]string, 0, len(dict))
	for _, group := range dict {
		result = append(result, group)
	}
	return result
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
