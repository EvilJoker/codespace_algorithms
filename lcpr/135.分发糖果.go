/*
 * @lc app=leetcode.cn id=135 lang=golang
 * @lcpr version=
 *
 * [135] 分发糖果
 */
package main

// @lc code=start
// 思路：
// 1. 每个孩子至少分到1个糖果
// 2. 相邻两个孩子中，评分高的必须获得更多糖果
// 3. 需要同时满足左右两个方向的规则，因此需要两次遍历
// 4. 第一次从左到右遍历，确保右边评分高的孩子获得更多糖果
// 5. 第二次从右到左遍历，确保左边评分高的孩子获得更多糖果
// 6. 取两次遍历的最大值作为最终糖果数

func candy(ratings []int) int {
	n := len(ratings)
	// 初始化糖果数组，每个孩子至少1个糖果
	candies := make([]int, n)
	candies[0] = 1

	// 从左到右遍历，处理右边评分高的情况
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}

	// 从右到左遍历，处理左边评分高的情况
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// 取两次遍历的最大值，确保同时满足左右规则
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// 计算总糖果数
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
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
