/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	helper(nums, []int{}, &ans)
	return ans
}
func helper(nums, trace []int, output *[][]int) {
	if len(nums) == 0 {
		ret := make([]int, len(trace))
		copy(ret, trace)
		*output = append(*output, ret)
		return
	}
	for i, n := range nums {
		trace = append(trace, n)
		nums[0], nums[i] = nums[i], nums[0]
		helper(nums[1:len(nums)], trace, output)
		nums[0], nums[i] = nums[i], nums[0]
		trace = trace[0 : len(trace)-1]
	}
}
