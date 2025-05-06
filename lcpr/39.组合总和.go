/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=
 *
 * [39] 组合总和
 */
package main

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	// 解题思路：
	// 1. 使用回溯算法，每个数字可以重复使用
	// 2. 使用DFS遍历所有可能的组合
	// 3. 使用start参数避免重复组合（如[2,3]和[3,2]）
	// 4. 当sum等于target时，将当前路径加入结果集
	// 5. 当sum大于target时，剪枝返回
	// 时间复杂度：O(2^n)，空间复杂度：O(n)

	result := [][]int{}
	n := len(candidates)

	// 定义DFS函数
	var dfs func(start int, path []int, sum int)
	dfs = func(start int, path []int, sum int) {
		// 找到目标组合
		if sum == target {
			// 创建path的副本并加入结果集
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 剪枝：当前和超过目标值
		if sum > target {
			return
		}

		// 遍历所有可能的候选数字
		for i := start; i < n; i++ {
			currentNum := candidates[i]
			// 继续搜索，注意start参数为i，允许重复使用当前数字
			dfs(i, append(path, currentNum), sum+currentNum)
		}
	}

	// 从第一个数字开始搜索
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
