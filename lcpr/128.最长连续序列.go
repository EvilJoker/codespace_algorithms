/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=
 *
 * [128] 最长连续序列
 */
package main

// @lc code=start
// 思路：
// 1. 使用哈希表记录所有数字，用于快速查找
// 2. 对于每个数字，如果它是连续序列的起点（即num-1不在哈希表中），则从它开始向后查找连续序列
// 3. 记录最长连续序列的长度
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用哈希表存储所有数字
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0

	// 遍历每个数字
	for num := range numSet {
		// 如果当前数字是连续序列的起点
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// 向后查找连续序列
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// 更新最长序列长度
			maxLength = max(maxLength, currentLength)
		}
	}

	return maxLength
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
