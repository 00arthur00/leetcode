/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */
//此处无重复元素
func searchInsert1(nums []int, target int) int {
	b := 0
	e := len(nums) - 1
	for b <= e {
		mid := (b + e) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			b = mid + 1
		} else if nums[mid] > target {
			e = mid - 1
		}
	}
	return b
}

//,若有重复元素,题目为寻找target的左边届,即小于target的元素的个数
func searchInsert(nums []int, target int) int {
	b := 0
	e := len(nums)
	for b < e {
		mid := (b + e) / 2
		if nums[mid] == target {
			e = mid
		} else if nums[mid] < target {
			b = mid + 1
		} else if nums[mid] > target {
			e = mid
		}
	}
	return b
}

