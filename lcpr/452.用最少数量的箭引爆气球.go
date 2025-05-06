/*
 * @lc app=leetcode.cn id=452 lang=golang
 * @lcpr version=
 *
 * [452] 用最少数量的箭引爆气球
 */
package main

import "sort"

// @lc code=start
/*
思路：
1. 贪心算法：将气球按照左边界排序，然后从左到右遍历
2. 维护一个当前箭可以射中的区间范围 [maxLeft, minRight]
3. 当遇到新的气球时：
   - 如果新气球的左边界 > 当前区间右边界，说明需要新的箭
   - 否则，更新当前区间范围，取交集
*/
func findMinArrowShots(points [][]int) int {
	// 处理空数组情况
	if len(points) == 0 {
		return 0
	}

	// 按左边界排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 初始化第一支箭的区间范围
	arrowCount := 1
	currentLeft := points[0][0]
	currentRight := points[0][1]

	// 遍历所有气球
	for i := 1; i < len(points); i++ {
		balloonLeft := points[i][0]
		balloonRight := points[i][1]

		// 当前气球与箭的区间无重叠，需要新的箭
		if balloonLeft > currentRight {
			arrowCount++
			currentLeft = balloonLeft
			currentRight = balloonRight
			continue
		}

		// 更新当前箭的区间范围（取交集）
		currentLeft = max(currentLeft, balloonLeft)
		currentRight = min(currentRight, balloonRight)
	}

	return arrowCount
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
