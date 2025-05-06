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
	/*
	   解题思路：
	   1. 贪心算法：尽可能多地放置单词，直到放不下为止
	   2. 对于每一行：
	      - 计算可以放置的单词数量
	      - 计算剩余空间，用于分配空格
	      - 如果是最后一行，则左对齐
	      - 否则，均匀分配空格（左侧空格数可以比右侧多）
	   3. 特殊情况处理：
	      - 只有一个单词时左对齐
	      - 最后一行左对齐
	*/

	start, end, leftWidth := 0, 0, maxWidth
	result := []string{}

	// 遍历所有单词
	for end < len(words) {
		// 如果当前单词放不下且不是第一个单词，则处理当前行
		if leftWidth < len(words[end]) && end > start {
			result = append(result, formatLine(words, start, end-1, maxWidth))
			leftWidth = maxWidth
			start = end
			continue
		}

		// 能放下当前单词，更新剩余宽度
		leftWidth -= (len(words[end]) + 1) // +1 是为了单词间的空格
		end++
	}

	// 处理最后一行（左对齐）
	if start < len(words) {
		result = append(result, formatLine(words, start, len(words)-1, maxWidth))
	}

	return result
}

// formatLine 格式化一行文本
func formatLine(words []string, start, end, maxWidth int) string {
	// 计算当前行的单词数
	wordCount := end - start + 1

	// 计算所有单词的总长度
	totalWordLength := 0
	for i := start; i <= end; i++ {
		totalWordLength += len(words[i])
	}

	// 计算需要分配的空格数
	totalSpaces := maxWidth - totalWordLength

	// 处理特殊情况：最后一行或只有一个单词时左对齐
	if end == len(words)-1 || wordCount == 1 {
		return strings.Join(words[start:end+1], " ") + strings.Repeat(" ", totalSpaces-wordCount+1)
	}

	// 计算基础空格数和额外空格数
	baseSpaces := totalSpaces / (wordCount - 1)  // 每个间隔的基础空格数
	extraSpaces := totalSpaces % (wordCount - 1) // 需要额外分配的空格数

	// 构建格式化后的行
	var result strings.Builder
	for i := start; i <= end; i++ {
		result.WriteString(words[i])

		// 最后一个单词后不需要加空格
		if i == end {
			continue
		}

		// 添加基础空格
		result.WriteString(strings.Repeat(" ", baseSpaces))

		// 分配额外空格（优先分配给左侧间隔）
		if extraSpaces > 0 {
			result.WriteString(" ")
			extraSpaces--
		}
	}

	return result.String()
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
