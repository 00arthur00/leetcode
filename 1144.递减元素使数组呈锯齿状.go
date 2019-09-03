/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 */
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func movesToMakeZigzag(nums []int) int {
	ans1, ans2 := 0, 0

	for index := 0; index < len(nums); index += 2 {
		if index == 0 {
			if nums[0] >= nums[1] {
				ans1 += nums[0] - nums[1] + 1
			}
		} else if index == len(nums)-1 {
			if nums[index] >= nums[index-1] {
				ans1 += nums[index] - nums[index-1] + 1
			}
		} else if nums[index] >= nums[index-1] || nums[index] >= nums[index+1] {
			ans1 += abs(nums[index]-min(nums[index-1], nums[index+1])) + 1
		}
	}

	for index := 1; index < len(nums); index += 2 {
		if index == len(nums)-1 {
			if nums[index] >= nums[index-1] {
				ans2 += nums[index] - nums[index-1] + 1
			}
		} else if nums[index] >= nums[index-1] || nums[index] >= nums[index+1] {
			ans2 += abs(nums[index]-min(nums[index-1], nums[index+1])) + 1
		}
	}

	return min(ans1, ans2)
}
