/*
 * @lc app=leetcode.cn id=452 lang=golang
 * @lcpr version=
 *
 * [452] 用最少数量的箭引爆气球
 */
package main

import "sort"

// @lc code=start
func findMinArrowShots(points [][]int) int {
	// sl：贪心，对区间排列后，从左到右移动剑，当脱离一个气球的范围时，那么就该增加一只箭了

	n := len(points)
	if n == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	count := 1
	max_left, min_right := points[0][0], points[0][1] // 最大的左右边界
	// 开始

	for i := 1; i < n; i++ {
		left, right := points[i][0], points[i][1]

		if left > min_right { // 超出之前的边界
			count++
			max_left = left
			min_right = right
			continue
		}

		max_left = max(max_left, left)
		min_right = min(min_right, right)
	}

	return count
}

// @lc code=end

/*
// @lcpr case=start
// [[10,16],[2,8],[1,6],[7,12]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[3,4],[5,6],[7,8]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2],[2,3],[3,4],[4,5]]\n
// @lcpr case=end

*/
