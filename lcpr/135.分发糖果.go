/*
 * @lc app=leetcode.cn id=135 lang=golang
 * @lcpr version=
 *
 * [135] 分发糖果
 */
package main

// @lc code=start
func candy(ratings []int) int {
	// 思路：左右两次遍历
	n := len(ratings)
	candys := make([]int, n)
	candys[0] = 1
	// 左
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} else {
			candys[i] = 1
		}
	}
	// 右
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candys[i] = max(candys[i], candys[i+1]+1) // 需要记忆
		}
	}
	sum := 0

	for _, v := range candys {
		sum += v
	}

	return sum
}

// @lc code=end

/*
// @lcpr case=start
// [1,0,2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

*/
