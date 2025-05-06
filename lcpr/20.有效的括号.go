/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=
 *
 * [20] 有效的括号
 */
package main

// @lc code=start
func isValid(s string) bool {
	// 解题思路：
	// 1. 使用栈来存储左括号，遇到右括号时与栈顶元素匹配
	// 2. 如果匹配成功，弹出栈顶元素；如果匹配失败，返回false
	// 3. 最后检查栈是否为空，为空则说明所有括号都匹配成功
	// 4. 时间复杂度：O(n)，空间复杂度：O(n)

	// 初始化栈
	stack := []byte{}

	// 遍历字符串
	for i := 0; i < len(s); i++ {
		c := s[i]

		// 如果是左括号，入栈
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			// 如果是右括号，且栈为空，说明没有匹配的左括号
			if len(stack) == 0 {
				return false
			}
			// 获取栈顶元素
			last := stack[len(stack)-1]
			// 检查括号是否匹配
			if c == ')' && last != '(' {
				return false
			}
			if c == ']' && last != '[' {
				return false
			}
			if c == '}' && last != '{' {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		}
	}

	// 检查栈是否为空，不为空说明有未匹配的左括号
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
