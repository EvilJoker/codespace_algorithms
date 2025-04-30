/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	//sl: 目前为止的和，如果比当前小，

	maxSum := nums[0]
	curSum := nums[0]

	for i := 1; i < len(nums); i++ {
		curSum = max(nums[i], curSum+nums[i]) // 如果之前的累加是负的就舍弃
		maxSum = max(curSum, maxSum)
	}

	return maxSum

}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
