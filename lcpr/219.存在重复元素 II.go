/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=
 *
 * [219] 存在重复元素 II
 */
package main

// @lc code=start
// 思路：使用哈希表记录每个数字最后一次出现的位置
// 遍历数组时，如果当前数字在哈希表中存在，且当前位置与上次出现位置的距离不超过k，则返回true
// 否则更新哈希表中该数字的位置为当前位置
func containsNearbyDuplicate(nums []int, k int) bool {
	// lastPos 记录每个数字最后一次出现的位置
	lastPos := map[int]int{}

	for i, num := range nums {
		// 获取当前数字上次出现的位置
		lastIndex := lastPos[num]
		// 如果数字之前出现过，且距离不超过k，返回true
		if lastIndex != 0 && i-lastIndex+1 <= k {
			return true
		}
		// 更新数字的最后出现位置
		lastPos[num] = i + 1
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
