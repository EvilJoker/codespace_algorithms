/**
 * 题目名称：跳跃游戏（Jump Game）
 * 
 * 题目理解：
 * 给定一个非负整数数组 nums，每个元素表示从当前位置可以向前跳跃的最大长度。目标是确定到达数组最后一个元素的最小跳跃次数。
 *
 * 解题思路：
 * 1. 初始化当前能达到的最远位置 `farthest` 为数组的第一个元素 `nums[0]`，跳跃次数 `steps` 为 0。
 * 2. 使用一个循环遍历数组，对于每个位置 i，如果它能直接或间接跳到 `farthest` 之后的位置，则更新 `farthest` 为可达的最远位置，并且每当 `farthest` 增加时，增加跳跃次数 `steps`。
 * 3. 当 `farthest` 达到最后一个元素时跳出循环，并返回跳跃次数 `steps`。
 *
 * 注意要点：
 * - 要特别注意边界条件，确保不会访问超出数组范围的元素。
 * - 每次更新 `farthest` 时需要判断能否达到新的未探索位置，而不是盲目跟随最大跳跃距离。
 * - 可以使用贪心策略来解决问题，每次尽可能选择能够使得跳跃更接近终点的位置。
 * 
 * 时间复杂度：
 * O(n)，其中 n 是数组 nums 的长度。因为我们只需要遍历一次数组即可找到最小跳跃次数。
 */

public int jump(int[] nums) {
    // 起始跳跃次数为 0
    int steps = 0;
    
    // 初始状态下能到达的最远索引
    int farthest = nums[0];
    
    // 当前遍历到的索引位置
    int current = 0;

    // 遍历数组，直到能够到达最后一个元素
    while (current < nums.length - 1) {
        // 更新当前能达到的最远位置
        for (int next = current + 1; next <= Math.min(nums.length - 1, current + nums[current]); next++) {
            // 如果下一次跳跃可以达到比之前更远的地方，则更新 farthest
            // 遍历每一步，尝试找出
            farthest = Math.max(nums[next] + next, farthest);

        }
        
        // 当前跳跃使我们到达了新的最远位置，增加跳跃次数
        steps++;
        
        // 更新当前遍历到的索引位置为当前能达到的最远位置
        current = farthest;
    }

    // 返回到达终点所需的最小跳跃次数
    return steps;
}