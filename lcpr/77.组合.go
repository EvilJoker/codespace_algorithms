/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	//sl: 组合是不管次序的，排列是管的。去除重复的办法是，遍历时不重复遍历

	result := [][]int{}

	var dfs func(index int, path []int)

	dfs = func(index int, path []int) {

		// 终止条件
		if len(path) == k {
			// 复制一份，再加入
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}

		// 可以优化
		for i := index; i <= n; i++ {
			dfs(i+1, append(path, i))
		}
	}

	dfs(1, []int{})

	return result

}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/

