/*
 * @lc app=leetcode.cn id=134 lang=golang
 * @lcpr version=
 *
 * [134] 加油站
 */
package main

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	// 解题思路：
	// 1. 如果总油量小于总消耗，则无法完成环形路线
	// 2. 从起点开始，如果当前累积油量小于消耗，则从下一个位置重新开始
	// 3. 如果存在可行解，则最后找到的起点就是答案

	start := 0      // 起始加油站位置
	totalGas := 0   // 总油量
	totalCost := 0  // 总消耗
	currentGas := 0 // 当前累积油量

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currentGas += gas[i] - cost[i]

		// 如果当前累积油量为负，说明从start到i的路线不可行
		// 将下一个位置作为新的起点
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}

	// 如果总油量小于总消耗，则无法完成环形路线
	if totalGas < totalCost {
		return -1
	}

	return start
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n[3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[3,4,3]\n
// @lcpr case=end

*/
