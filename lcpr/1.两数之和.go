/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	/* 思路： 使用hash
	 */
	dict := make(map[int]int, 0)

	for i, e := range nums {
		another := target - e
		if index, ok := dict[another]; ok {
			return []int{index, i}
		}
		dict[e] = i
	}

	return []int{0, 0}

}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [3,2,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [3,3]\n6\n
// @lcpr case=end

*/
