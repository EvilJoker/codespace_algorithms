/*
 * @lc app=leetcode.cn id=274 lang=golang
 * @lcpr version=
 *
 * [274] H 指数
 */
package main

import "sort"

// @lc code=start
func hIndex(citations []int) int {
	// 排序，如果 某个索引的序号小于 论文数，就找到了h
	sort.Ints(citations)
	n := len(citations)

	for i := 0; i < n; i++ {
		if citations[n-1-i] < i+1 {
			return i
		}
	}
	return n
}

// @lc code=end

/*
// @lcpr case=start
// [3,0,6,1,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,3,1]\n
// @lcpr case=end

*/
