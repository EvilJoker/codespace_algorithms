/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if c == ')' && last != '(' {
				return false
			}
			if c == ']' && last != '[' {
				return false
			}
			if c == '}' && last != '{' {
				return false
			}
			stack = stack[:len(stack)-1]

		}
	}

	if len(stack) != 0 {
		return false
	}

	return true

}

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

*/
