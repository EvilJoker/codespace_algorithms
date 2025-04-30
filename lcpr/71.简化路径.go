/*
 * @lc app=leetcode.cn id=71 lang=golang
 * @lcpr version=
 *
 * [71] 简化路径
 */
package main

import "strings"

// @lc code=start
func simplifyPath(path string) string {
	//sl: 栈
	strs := strings.Split(path, "/")
	stack := []string{}

	for _, str := range strs {

		if str == "" {
			continue
		}
		if str == "." {
			continue
		}
		if str == ".." {

			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, str)

	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end

/*
// @lcpr case=start
// "/home/"\n
// @lcpr case=end

// @lcpr case=start
// "/home//foo/"\n
// @lcpr case=end

// @lcpr case=start
// "/home/user/Documents/../Pictures"\n
// @lcpr case=end

// @lcpr case=start
// "/../"\n
// @lcpr case=end

// @lcpr case=start
// "/.../a/../b/c/../d/./"\n
// @lcpr case=end

*/
