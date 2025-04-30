/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=
 *
 * [6] Z 字形变换
 */
package main

// @lc code=start
func convert(s string, numRows int) string {
	//sl: 记录方向
	// 向下 和向右切换

	if numRows == 1 {
		return s
	}

	direction := 0 // 0下， 1 右
	n := len(s)

	strs := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		strs[i] = make([]byte, n)
	}
	p, q := 0, 0

	for i := 0; i < n; {
		if direction == 0 {

			strs[p][q] = s[i]
			if p == numRows-1 { // 等于下边界
				direction = 1
				continue
			}
			p++

		} else {
			strs[p][q] = s[i]

			if p == 0 {
				direction = 0
				continue
			}
			p--
			q++

		}
		i++
	}
	res := ""

	// 解析
	for _, str := range strs {
		for _, b := range str {
			if b != 0 {
				res += string(b)
			}
		}
	}
	return res

}

// @lc code=end

/*


// @lcpr case=start
// "A"\n1\n
// @lcpr case=end

*/
