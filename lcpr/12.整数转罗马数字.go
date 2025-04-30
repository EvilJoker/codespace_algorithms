/*
 * @lc app=leetcode.cn id=12 lang=golang
 * @lcpr version=
 *
 * [12] 整数转罗马数字
 */
package main

// @lc code=start

import "sort"

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
	// 思路： 贪心，不断取最大

	keys := []int{}
	for key := range dict1 {
		keys = append(keys, key)
	}
	sort.Ints(keys) // 从小到大
	n := len(keys)

	res := ""
	for num > 0 {
		for i := n - 1; i >= 0; {
			val := keys[i]
			str, _ := dict1[val]
			if num >= val {
				res += str
				num -= val
			} else {
				i--
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
