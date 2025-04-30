/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=
 *
 * [22] 括号生成
 */
package main

// @lc code=start
func generateParenthesis(n int) []string {
	//sj: 右括号，不能超过左括号。独立数字记录 左括号数量

	result := []string{}

	var dfs func(path string, left int, right int) // 左右括号可用数
	dfs = func(path string, left int, right int) {
		if left == 0 && right == 0 {
			result = append(result, path)
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if right < left { // 右括号比左括号用的多
			return
		}

		dfs(path+"(", left-1, right)

		dfs(path+")", left, right-1)

	}
	dfs("", n, n)
	return result
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
