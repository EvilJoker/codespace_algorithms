/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=
 *
 * [45] 跳跃游戏 II
 */
package main

// @lc code=start
func jump(nums []int) int {
	// 解题思路：
	// 1. 使用贪心算法，从后向前遍历
	// 2. 对于每个位置，找到能够到达当前位置的最远位置
	// 3. 每次选择能够到达当前位置的最远位置作为新的目标位置
	// 4. 重复此过程直到到达数组起始位置
	// 时间复杂度：O(n²)，空间复杂度：O(1)

	n := len(nums)
	if n <= 1 {
		return 0
	}

	steps := 0
	targetPosition := n - 1 // 目标位置初始化为数组末尾

	// 当目标位置不是起始位置时继续循环
	for targetPosition > 0 {
		// 从前往后遍历，找到能够到达目标位置的最远位置
		for i := 0; i <= targetPosition; i++ {
			// 如果当前位置能够到达目标位置
			if i+nums[i] >= targetPosition {
				targetPosition = i // 更新目标位置
				steps++            // 步数加1
				break              // 找到最远位置后退出内层循环
			}
		}
	}

	return steps
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,0,1,4]\n
// @lcpr case=end

*/
