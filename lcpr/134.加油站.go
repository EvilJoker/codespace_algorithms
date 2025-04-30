/*
 * @lc app=leetcode.cn id=134 lang=golang
 * @lcpr version=
 *
 * [134] 加油站
 */
package main

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	//思路： 总的 cost 大于 总 gas 就是达不到。当完不成时，切换下一个节点为新起点
	start, sumGas, sumCost := 0, 0, 0

	for i := 0; i < len(cost); i++ {
		sumGas += gas[i]
		sumCost += cost[i]
		if sumCost > sumGas {
			start = i
		}
	}
	if sumCost > sumGas {
		return -1
	}

	return start

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n[3,4,5,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n[3,4,3]\n
// @lcpr case=end

*/
