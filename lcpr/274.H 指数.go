/*
 * @lc app=leetcode.cn id=274 lang=golang
 * @lcpr version=
 *
 * [274] H 指数
 */
package main

import "sort"

// @lc code=start
/*
思路：
1. H指数的定义：至少有h篇论文被引用了h次或以上
2. 将引用次数排序后，从大到小遍历
3. 对于每个位置i，检查是否有i+1篇论文的引用次数大于等于i+1
4. 如果找到第一个不满足条件的位置，返回i
5. 如果全部满足，返回论文总数n
*/
func hIndex(citations []int) int {
	// 将引用次数从小到大排序
	sort.Ints(citations)
	n := len(citations)

	// 从大到小遍历，检查每个位置是否满足H指数的条件
	for i := 0; i < n; i++ {
		// 如果第i+1大的引用次数小于i+1，说明找到了H指数
		if citations[n-1-i] < i+1 {
			return i
		}
	}
	// 如果所有论文都满足条件，返回论文总数
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
