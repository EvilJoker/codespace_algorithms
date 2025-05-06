/*
 * @lc app=leetcode.cn id=12 lang=golang
 * @lcpr version=
 *
 * [12] 整数转罗马数字
 */
package main

// @lc code=start

var dict1 map[int]string = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) string {
	// 思路：贪心算法，从大到小遍历罗马数字对应的值
	// 时间复杂度：O(1)，空间复杂度：O(1)

	// 预定义罗马数字对应的值数组，按从大到小排序
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	res := ""
	for num > 0 {
		for _, val := range values {
			if num >= val {
				res += dict1[val]
				num -= val
				break // 找到当前最大可用的值后，重新开始遍历
			}
		}
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// 3749\n
// @lcpr case=end

// @lcpr case=start
// 58\n
// @lcpr case=end

// @lcpr case=start
// 1994\n
// @lcpr case=end

*/
