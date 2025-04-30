/*
 * @lc app=leetcode.cn id=57 lang=golang
 * @lcpr version=
 *
 * [57] 插入区间
 */
package main

import "sort"

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	// 思路： 找到 区间改在的位置，新区间起点 大于 上一个区间的起点，小于 下一个区间的起点
	// 1. 从位置 上一个区间开始合并
	// 2. 2总关系： 上一个对下一个包含； 上一个对下一个重叠； 一直遍历结尾

	// 加入然后排序
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 从前到后合

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		cur := intervals[i]
		// 被前者包含
		if cur[1] <= last[1] {
			continue
		}
		// 和前者重叠
		if cur[0] <= last[1] && cur[1] >= last[1] {
			last = []int{last[0], cur[1]}
			result = append(result[:len(result)-1], last)
			continue
		}
		if cur[0] > last[1] {
			// 独立
			result = append(result, cur)
		}

	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[6,9]]\n[2,5]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,5],[6,7],[8,10],[12,16]]\n[4,8]\n
// @lcpr case=end

*/
