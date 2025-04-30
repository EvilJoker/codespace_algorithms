/*
 * @lc app=leetcode.cn id=56 lang=golang
 * @lcpr version=
 *
 * [56] 合并区间
 */
package main

import "sort"

// @lc code=start
func merge(intervals [][]int) [][]int {
	// 思路：先排序后合并

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//合并
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		cur := intervals[i]
		// 包含
		if last[1] >= cur[1] {
			continue
		}
		// 重叠
		if cur[0] <= last[1] && cur[1] >= last[1] {
			res = res[:len(res)-1]
			res = append(res, []int{last[0], cur[1]})
			continue
		}
		// 独立
		res = append(res, cur)

	}

	return res

}

// @lc code=end

/*
// @lcpr case=start
// [[1,3],[2,6],[8,10],[15,18]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,4],[4,5]]\n
// @lcpr case=end

*/
