/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=
 *
 * [55] 跳跃游戏
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	// 解题思路：
	// 1. 使用贪心算法，维护一个变量记录能够到达的最远位置
	// 2. 遍历数组，对于每个位置i：
	//    - 如果i超过了最远位置，说明无法到达该位置，返回false
	//    - 否则，更新最远位置为max(当前最远位置, i + nums[i])
	// 3. 如果最远位置能够到达或超过最后一个位置，返回true
	// 时间复杂度：O(n)，空间复杂度：O(1)

	// 记录能够到达的最远位置
	maxReachable := 0

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 如果当前位置超过了最远可达位置，说明无法到达
		if i > maxReachable {
			return false
		}

		// 计算从当前位置能够到达的最远位置
		currentReachable := i + nums[i]

		// 如果能够到达或超过最后一个位置，直接返回true
		if currentReachable >= len(nums)-1 {
			return true
		}

		// 更新最远可达位置
		maxReachable = max(maxReachable, currentReachable)
	}

	// 检查最终的最远位置是否能够到达最后一个位置
	return maxReachable >= len(nums)-1
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
