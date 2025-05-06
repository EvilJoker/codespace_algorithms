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
	// 解题思路：
	// 1. 首先按照区间的起始位置对所有区间进行排序
	// 2. 遍历排序后的区间，对于每个区间，有三种情况：
	//    - 当前区间被前一个区间完全包含
	//    - 当前区间与前一个区间有重叠
	//    - 当前区间与前一个区间完全独立
	// 3. 根据不同的情况，更新结果数组
	// 时间复杂度：O(nlogn)，主要来自排序
	// 空间复杂度：O(n)，需要存储结果数组

	// 特殊情况处理
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果数组，将第一个区间加入
	result := [][]int{intervals[0]}

	// 遍历剩余区间
	for i := 1; i < len(intervals); i++ {
		lastInterval := result[len(result)-1] // 结果数组中的最后一个区间
		currentInterval := intervals[i]       // 当前待处理的区间

		// 情况1：当前区间被前一个区间完全包含
		if lastInterval[1] >= currentInterval[1] {
			continue
		}

		// 情况2：当前区间与前一个区间有重叠
		if currentInterval[0] <= lastInterval[1] {
			// 更新前一个区间的结束位置
			result[len(result)-1][1] = currentInterval[1]
			continue
		}

		// 情况3：当前区间与前一个区间完全独立
		result = append(result, currentInterval)
	}

	return result
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
