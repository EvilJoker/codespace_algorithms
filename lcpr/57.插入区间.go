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
	/*
	   解题思路：
	   1. 将新区间加入原区间数组，并按区间起始点排序
	   2. 遍历排序后的区间数组，处理区间合并的三种情况：
	      - 当前区间被前一个区间完全包含
	      - 当前区间与前一个区间有重叠
	      - 当前区间与前一个区间完全独立
	   3. 返回合并后的区间数组
	*/

	// 将新区间加入原数组并排序
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组，将第一个区间加入
	result := [][]int{intervals[0]}

	// 遍历剩余区间进行合并
	for i := 1; i < len(intervals); i++ {
		lastInterval := result[len(result)-1] // 结果数组中的最后一个区间
		currentInterval := intervals[i]       // 当前待处理的区间

		// 情况1：当前区间被前一个区间完全包含
		if currentInterval[1] <= lastInterval[1] {
			continue
		}

		// 情况2：当前区间与前一个区间有重叠
		if currentInterval[0] <= lastInterval[1] && currentInterval[1] >= lastInterval[1] {
			// 更新前一个区间的结束位置
			lastInterval = []int{lastInterval[0], currentInterval[1]}
			result[len(result)-1] = lastInterval
			continue
		}

		// 情况3：当前区间与前一个区间完全独立
		if currentInterval[0] > lastInterval[1] {
			result = append(result, currentInterval)
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
