/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=
 *
 * [202] 快乐数
 */
package main

// @lc code=start

/*
解题思路：
1. 快乐数的定义：对于一个正整数，每次将该数替换为它每个位置上的数字的平方和，重复这个过程直到这个数变为1
2. 如果这个数最终变为1，则是快乐数；如果陷入循环，则不是快乐数
3. 使用哈希表记录已经出现过的数字，如果重复出现，说明陷入了循环
*/

func isHappy(n int) bool {
	// visited 记录已经出现过的数字
	visited := map[int]bool{}

	// 当n不等于1时继续循环
	for n != 1 {
		// 如果数字已经出现过，说明陷入了循环，返回false
		if visited[n] {
			return false
		}
		// 标记当前数字已经访问过
		visited[n] = true
		// 计算下一个数字
		n = calculateNext(n)
	}

	return true
}

// calculateNext 计算下一个数字
// 将数字的每一位平方后求和
func calculateNext(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10 // 获取最后一位数字
		n /= 10         // 去掉最后一位
		sum += digit * digit
	}
	return sum
}

// @lc code=end

/*
// @lcpr case=start
// 19\n
// @lcpr case=end

// @lcpr case=start
// 2\n
// @lcpr case=end

*/
