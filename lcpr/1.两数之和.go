/*
 * @lc app=leetcode.cn id=1 lang=golang
 * @lcpr version=
 *
 * [1] 两数之和
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	/*
		解题思路：哈希表
		1. 使用哈希表存储已经遍历过的数字及其索引
		2. 遍历数组，对于每个数字nums[i]：
		   - 计算目标值target与当前数字的差值another = target - nums[i]
		   - 检查差值another是否在哈希表中
		   - 如果存在，说明找到了两个数之和等于target，返回它们的索引
		   - 如果不存在，将当前数字及其索引存入哈希表
		3. 时间复杂度：O(n)，空间复杂度：O(n)
	*/
	// 创建哈希表，key为数字，value为索引
	dict := make(map[int]int)

	// 遍历数组
	for i, num := range nums {
		// 计算目标值与当前数字的差值
		another := target - num
		// 检查差值是否在哈希表中
		if index, ok := dict[another]; ok {
			// 找到答案，返回两个数的索引
			return []int{index, i}
		}
		// 将当前数字及其索引存入哈希表
		dict[num] = i
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
