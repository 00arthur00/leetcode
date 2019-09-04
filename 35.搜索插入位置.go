/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
func searchInsert(nums []int, target int) int {
	b := 0
	e := len(nums) - 1
	for {
		if b > e {
			return b
		}
		mid := (b + e) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			b = mid + 1
		} else {
			e = mid - 1
		}
	}
}

