/*
 * @lc app=leetcode.cn id=3 lang=java
 *
 * [3] 无重复字符的最长子串
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 
 */

// @lc code=start

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.List;

class Solution {
    public int lengthOfLongestSubstring(String s) {
        // 滑动窗口
        // 维护 left 和 right 两个指针, 0
        // rigth 不断向右滑动，不断更新最大长度
        // 当出现重复字符时，left 不断向右移动，直到 不重复为止

        int left = 0;
        int right = 0;
        int maxlen = 0;

        ArrayList<Character> norepeat = new ArrayList<>();
        while (right < s.length()) {
            // right 有移
            char c = s.charAt(right);
            if (!norepeat.contains(c)) {
                right++;
                norepeat.add(c);
                // 更新当前最长
                maxlen = Math.max(maxlen, right - left);
                continue;
            }
            // 包含时left 左移
            while (norepeat.contains(c)) {
                char t = s.charAt(left);
                norepeat.remove(new Character(t));
                left++;
            }

        }
        return maxlen;

    }

    public static void main(String[] args) {
        new Solution().lengthOfLongestSubstring(" ");

    }
}
// @lc code=end
