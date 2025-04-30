/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=
 *
 * [39] 组合总和
 */
package main

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	//sl： 回溯，每一个数字可以无限用
	result := [][]int{}
	n := len(candidates)

	var dfs func(start int, path []int, sum int)
	dfs = func(start int, path []int, sum int) {
		// 终止条件
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		if sum > target {
			return
		}

		// 先找出可选的
		for i := start; i < n; i++ {
			can := candidates[i]
			// 数字不重复考虑，比如用完2 ，就不考虑2了。这里就是 i 的原因
			dfs(i, append(path, can), sum+can)
			// 回溯
		}

	}
	dfs(0, []int{}, 0)
	return result

}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
