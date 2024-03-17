/**
 * 题目名称：接雨水问题（Trapping Rain Water）
 *
 * 题目理解：
 *   给定一个非负整数数组，它表示每个位置的柱子高度，计算这些柱子能接住多少单位体积的雨水。假设每相邻两个柱子之间形成一个“槽”，雨水可以完全填充宽度为1的槽。
 *  https://zhuanlan.zhihu.com/p/107792266
 * 解题思路：
 *   1. 一个柱子能接的水，取决于 左边 和右边最高的柱子。相当于木桶的两边，取最低的那一边。 min(l_max_height ,r_max_hright) - height
 *   1. 双指针法。使用两个指针 left 和 right 分别从数组两端开始遍历.
     2. l_max ，指 0 ~ left 的最大高度，r_max 指 right ~ n 的最大高度
     3. 一般情况下，要找两边最高的柱子。但是实际上，只要知道，l_max 和 r_max 比较矮的一边即可。因为第一步 min(l_max_height ,r_max_hright)
     4. 当l_max < r_max ，即使 r_max 不是右边最高的柱子。有 l_max 就够了
 * 注意要点：
 *   - 要注意处理边界条件，例如当数组只有一个元素或为空时，雨水量为0。
 * 

 * 时间复杂度：
 *   - 由于我们需要遍历整个数组一次，并且每次操作的时间复杂度为 O(1)，所以总时间复杂度为 O(n)。

 */

 public class TrappingRainWater {
    public static int trap(int[] height) {
        if (height == null || height.length <= 1) {
            return 0; // 边界条件，数组长度为0或1时，没有雨水可接
        }

        int left = 0, right = height.length - 1;
        int result = 0;
        int l_max = height[left];
        int r_max = height[right];

        while (left <= right) {
            // 更新最大高度
            l_max = Math.max(l_max, height[left]); // 0 ~ left 最大高度
            r_max = Math.max(r_max, height[right]); // right ~ height.length 最大高度

            if (l_max < r_max) {
                // 左边的矮柱子找到了，虽然 r_max 不一定是右边最大，但是可以确定了
                result += l_max - height[left];
                left++;
            } else {
                // 右边的矮柱子找到了，虽然 l_max 不一定是左边最大，但是可以确定了
                result += r_max - height[right];
                right--;
            }
            
        }

        return result;
    }

    public static void main(String[] args) {
        int[] height = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
        System.out.println("能接住的雨水量: " + trap(height));
    }
}