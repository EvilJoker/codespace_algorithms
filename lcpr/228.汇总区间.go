/*
 * @lc app=leetcode.cn id=228 lang=golang
 * @lcpr version=
 *
 * [228] 汇总区间
 */
package main

import "strconv"

// @lc code=start
// 思路：
// 1. 遍历数组，维护当前区间的起始值 start 和期望的下一个值 end
// 2. 当实际值不等于期望值时，说明区间断开，需要记录当前区间并开始新区间
// 3. 最后需要处理最后一个区间
// 4. 将区间转换为要求的字符串格式

func summaryRanges(nums []int) []string {
	// 处理空数组的情况
	if len(nums) == 0 {
		return []string{}
	}

	// 存储所有区间
	ranges := [][]int{}

	// 初始化第一个区间的起始值
	start, expectedNext := nums[0], nums[0]

	// 遍历数组寻找连续区间
	for i := 0; i < len(nums); i++ {
		if expectedNext != nums[i] {
			// 区间断开，记录当前区间
			ranges = append(ranges, []int{start, expectedNext - 1})
			// 开始新区间
			start, expectedNext = nums[i], nums[i]
		}
		expectedNext++
	}

	// 处理最后一个区间
	ranges = append(ranges, []int{start, expectedNext - 1})

	// 将区间转换为字符串格式
	result := []string{}
	for _, r := range ranges {
		if r[0] == r[1] {
			// 单个数字的情况
			result = append(result, strconv.Itoa(r[0]))
		} else {
			// 区间的情况
			result = append(result, strconv.Itoa(r[0])+"->"+strconv.Itoa(r[1]))
		}
	}

	return result
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,2,4,5,7]\n
// @lcpr case=end

// @lcpr case=start
// [0,2,3,4,6,8,9]\n
// @lcpr case=end

*/
