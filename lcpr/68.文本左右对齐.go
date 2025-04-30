/*
 * @lc app=leetcode.cn id=68 lang=golang
 * @lcpr version=
 *
 * [68] 文本左右对齐
 */
package main

import "strings"

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	// sl: 先尽可能塞单词，塞不下时，计算间隔。
	start, end, leftWidth := 0, 0, maxWidth
	res := []string{}
	for end < len(words) {
		if leftWidth < len(words[end]) && end > start {
			// 塞不下，切换新一行
			res = append(res, fullblack(words, start, end-1, maxWidth))
			leftWidth = maxWidth
			start = end
			continue
		}

		// 能塞下
		leftWidth -= (len(words[end]) + 1)

		end++

	}

	// 处理最后一行（左对齐）
	if start < len(words) {
		res = append(res, fullblack(words, start, len(words)-1, maxWidth))
	}

	return res

}

func fullblack(words []string, start, end, maxWidth int) string {

	n := end - start + 1

	leftWidths := maxWidth
	for i := start; i <= end; i++ {
		leftWidths -= len(words[i])
	}

	// 最后一行|| 只有一个单词时

	if end == len(words)-1 || n == 1 {
		return strings.Join(words[start:end+1], " ") + strings.Repeat(" ", leftWidths-n+1)
	}

	basespaces := leftWidths / (n - 1)
	headspaces := leftWidths % (n - 1)

	// 多个单词，均匀分布，左比右舵
	res := ""
	for i := start; i <= end; i++ {

		res += words[i]
		if i == end {
			continue
		}
		res += strings.Repeat(" ", basespaces)
		if headspaces > 0 {
			res += " "
			headspaces -= 1
		}

	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// ["This", "is", "an", "example", "of", "text", "justification."]\n16\n
// @lcpr case=end

// @lcpr case=start
// ["What","must","be","acknowledgment","shall","be"]\n16\n
// @lcpr case=end


*/
