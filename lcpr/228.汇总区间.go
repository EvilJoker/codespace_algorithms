/*
 * @lc app=leetcode.cn id=228 lang=golang
 * @lcpr version=
 *
 * [228] 汇总区间
 */
package main

import "strconv"

// @lc code=start
func summaryRanges(nums []int) []string {
	/*
		sl: 从每个元素出发，不断累加。如果断档就生成一个新元素。对访问过的元素使用 bool 标记
	*/

	n := len(nums)

	if n == 0 {
		return []string{}
	}
	res := [][]int{}

	start, end := nums[0], nums[0]
	for i := 0; i < n; i++ {
		if end != nums[i] {
			res = append(res, []int{start, end - 1})
			start, end = nums[i], nums[i]
		}
		end++
	}
	// 最后一步

	res = append(res, []int{start, end - 1})

	res_strs := []string{}
	for _, v := range res {
		if v[0] != v[1] {
			res_strs = append(res_strs, strconv.Itoa(v[0])+"->"+strconv.Itoa(v[1]))
		} else {
			res_strs = append(res_strs, strconv.Itoa(v[0]))
		}

	}
	return res_strs

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
