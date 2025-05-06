/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=
 *
 * [77] 组合
 */
package main

// @lc code=start
func combine(n int, k int) [][]int {
	/*
		解题思路：回溯算法
		1. 问题分析：
		   - 从1到n中选择k个数的所有组合
		   - 每个组合中的数字不重复
		   - 组合顺序不重要，[1,2]和[2,1]视为相同组合

		2. 算法设计：
		   - 使用回溯法，通过DFS构建所有可能的组合
		   - 使用index参数控制选择范围，避免重复组合
		   - 当path长度等于k时，将当前组合加入结果集

		3. 复杂度分析：
		   - 时间复杂度：O(C(n,k))，即组合数
		   - 空间复杂度：O(k)，递归调用栈的深度

		4. 优化点：
		   - 剪枝：当剩余可选数字不足以填满k个位置时提前返回
		   - 使用index控制选择范围，避免重复组合
	*/

	// 存储所有组合的结果
	result := [][]int{}

	// 定义DFS函数，用于生成组合
	// index: 当前可选的起始数字
	// path: 当前已选择的数字路径
	var dfs func(index int, path []int)

	dfs = func(index int, path []int) {
		// 当路径长度达到k时，找到一个有效组合
		if len(path) == k {
			// 创建当前路径的副本并加入结果集
			comb := make([]int, k)
			copy(comb, path)
			result = append(result, comb)
			return
		}

		// 从index开始遍历可选数字
		// 剪枝优化：当剩余数字不足以填满k个位置时，可以提前结束
		for i := index; i <= n; i++ {
			// 选择当前数字，继续DFS
			dfs(i+1, append(path, i))
		}
	}

	// 从1开始DFS搜索
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
