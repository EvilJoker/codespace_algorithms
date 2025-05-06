/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=
 *
 * [6] Z 字形变换
 */
package main

// @lc code=start
func convert(s string, numRows int) string {
	/*
		解题思路：模拟Z字形变换过程
		1. 使用二维数组存储字符，按Z字形规则填充
		2. 使用direction变量控制移动方向：0表示向下，1表示向右上
		3. 当到达边界时切换方向
		4. 最后按行遍历二维数组，收集非空字符
		时间复杂度：O(n)，空间复杂度：O(n*numRows)
	*/

	// 特殊情况：只有一行时直接返回原字符串
	if numRows == 1 {
		return s
	}

	direction := 0 // 0表示向下移动，1表示向右上移动
	n := len(s)

	// 初始化二维数组，用于存储Z字形排列的字符
	strs := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		strs[i] = make([]byte, n)
	}

	// p表示当前行，q表示当前列
	p, q := 0, 0

	// 遍历字符串，按Z字形规则填充二维数组
	for i := 0; i < n; {
		if direction == 0 {
			// 向下移动
			strs[p][q] = s[i]
			if p == numRows-1 { // 到达下边界，切换方向
				direction = 1
				continue
			}
			p++
		} else {
			// 向右上移动
			strs[p][q] = s[i]
			if p == 0 { // 到达上边界，切换方向
				direction = 0
				continue
			}
			p--
			q++
		}
		i++
	}

	// 按行遍历二维数组，收集非空字符
	res := ""
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
