/*
 * @lc app=leetcode.cn id=120 lang=golang
 * @lcpr version=30104
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode.cn/problems/triangle/description/
 *
 * algorithms
 * Medium (69.35%)
 * Likes:    1440
 * Dislikes: 0
 * Total Accepted:    417.4K
 * Total Submissions: 601.9K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形 triangle ，找出自顶向下的最小路径和。
 *
 * 每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1
 * 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
 * 输出：11
 * 解释：如下面简图所示：
 * ⁠  2
 * ⁠ 3 4
 * ⁠6 5 7
 * 4 1 8 3
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 *
 * 示例 2：
 *
 * 输入：triangle = [[-10]]
 * 输出：-10
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= triangle.length <= 200
 * triangle[0].length == 1
 * triangle[i].length == triangle[i - 1].length + 1
 * -10^4 <= triangle[i][j] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？
 *
 *
 */
package main

// @lc code=start
/* 解题思路：
1. 动态规划解法：
   - 定义dp数组：dp[i][j]表示从三角形顶部到位置(i,j)的最小路径和
   - 初始化：dp[0][0] = triangle[0][0]
   - 状态转移：
     - 对于每行的第一个元素：dp[i][0] = dp[i-1][0] + triangle[i][0]
     - 对于每行的最后一个元素：dp[i][i] = dp[i-1][i-1] + triangle[i][i]
     - 对于中间元素：dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
   - 最终结果：返回最后一行的最小值

*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// 初始化dp数组
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	// 初始化第一行
	dp[0][0] = triangle[0][0]

	// 动态规划填充dp数组
	for i := 1; i < n; i++ {
		// 处理每行的第一个元素
		dp[i][0] = dp[i-1][0] + triangle[i][0]

		// 处理中间元素
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}

		// 处理每行的最后一个元素
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	// 找出最后一行的最小值
	minSum := dp[n-1][0]
	for j := 1; j < n; j++ {
		minSum = min(minSum, dp[n-1][j])
	}

	return minSum
}

// @lc code=end

/*
// @lcpr case=start
// [[2],[3,4],[6,5,7],[4,1,8,3]]\n
// @lcpr case=end

// @lcpr case=start
// [[-10]]\n
// @lcpr case=end

*/
