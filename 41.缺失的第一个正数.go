/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num >= 1 && num <= len(nums) && num != nums[num-1] {
			nums[num-1], nums[i] = num, nums[num-1]
			i--
		}
	}
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}

