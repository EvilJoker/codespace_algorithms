/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=
 *
 * [219] 存在重复元素 II
 */
package main

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}

	for i, v := range nums {
		index := dict[v]
		if index != 0 && i-index+1 <= k {
			return true
		}

		// 规避 0值
		dict[v] = i + 1
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/
