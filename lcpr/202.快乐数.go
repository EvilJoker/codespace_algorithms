/*
 * @lc app=leetcode.cn id=202 lang=golang
 * @lcpr version=
 *
 * [202] 快乐数
 */
package main

// @lc code=start
func isHappy(n int) bool {
	//sl: 如果出现循环，就说明到不了1
	dict := map[int]bool{}

	for n != 1 {
		if _, exist := dict[n]; exist {
			return false
		}
		dict[n] = true
		n = getNext(n)
	}

	return true

}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		end := n % 10 // 最后一位
		n /= 10       // 下一位
		sum += end * end
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
