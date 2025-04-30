/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=
 *
 * [55] 跳跃游戏
 */
package main

// @lc code=start
func canJump(nums []int) bool {
	// 思路：维护一个最远位置，到不了就是异常

	remote := 0

	for i := 0; i < len(nums); i++ {
		if i > remote {
			return false
		}
		max_remote := i + nums[i]
		if max_remote >= len(nums)-1 {
			return true
		}

		remote = max(remote, max_remote)
	}

	if remote >= len(nums)-1 {
		return true
	} else {
		return false
	}

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
