/*
 * @lc app=leetcode.cn id=392 lang=golang
 * @lcpr version=
 *
 * [392] 判断子序列
 */
package main

// @lc code=start
func isSubsequence(s string, t string) bool {
	// 思路：双指针法
	// 1. 使用两个指针p和q分别指向s和t的起始位置
	// 2. 如果t[q] == s[p]，说明找到了一个匹配字符，p右移
	// 3. 无论是否匹配，q都要右移继续寻找下一个可能的匹配
	// 4. 如果p到达s末尾，说明s是t的子序列
	// 5. 如果q到达t末尾但p未到达s末尾，说明s不是t的子序列

	// 获取字符串长度
	sLen, tLen := len(s), len(t)
	// 如果s为空，则一定是t的子序列
	if sLen == 0 {
		return true
	}

	// 初始化双指针
	sPtr, tPtr := 0, 0

	// 遍历t字符串
	for tPtr < tLen {
		// 找到匹配字符
		if t[tPtr] == s[sPtr] {
			sPtr++
			// 如果s的所有字符都匹配成功
			if sPtr == sLen {
				return true
			}
		}
		tPtr++
	}

	return false
}

// @lc code=end

/*
// @lcpr case=start
// "abc"\n"ahbgdc"\n
// @lcpr case=end

// @lcpr case=start
// "axc"\n"ahbgdc"\n
// @lcpr case=end

*/
