/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=
 *
 * [128] 最长连续序列
 */
package main

import (
	"math"
	"sort"
)

// @lc code=start
func longestConsecutive(nums []int) int {
	//sl: 先排序，然后 从最小值，开始累加。利用hash 判断存不存在，不存在就 +1

	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	visted := map[int]bool{}
	for _, v := range nums {
		visted[v] = false
	}

	n, maxlen := len(nums), math.MinInt32

	for i := 0; i < n; i++ {
		cur := nums[i]
		count := 1
		for {
			v, exist := visted[cur]

			// 不存在
			if !exist {
				break
			}
			// 访问过
			if v {
				break
			}
			maxlen = max(maxlen, count)
			count++

			visted[cur] = true // 标识访问过
			cur++
		}
	}

	return maxlen

}

// @lc code=end

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,2]\n
// @lcpr case=end

*/
