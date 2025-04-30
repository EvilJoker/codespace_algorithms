/*
 * @lc app=leetcode.cn id=45 lang=golang
 * @lcpr version=
 *
 * [45] 跳跃游戏 II
 */
package main

// @lc code=start
func jump(nums []int) int {

	// 思路: 使用贪心，倒序访问，每次尽力向前跳
	n, step := len(nums), 0

	position := n - 1

	for position > 0 {
		for i := 0; i <= position; i++ {
			if i+nums[i] >= position {
				position = i
				step++
				break
			}
		}
	}

	return step

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
