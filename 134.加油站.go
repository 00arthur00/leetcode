/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) != len(cost) {
		return -1
	}
	var total_trunk, curr_trunk, start_station int
	for i := 0; i < len(gas); i++ {
		total_trunk += gas[i] - cost[i]
		curr_trunk += gas[i] - cost[i]
		if curr_trunk < 0 {
			start_station = i + 1
			curr_trunk = 0
		}
	}
	if total_trunk < 0 {
		return -1
	}
	return start_station
}

// @lc code=end

