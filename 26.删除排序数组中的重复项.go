/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//两个指针:遍历指针和重复指针,不相等的时候覆盖重复指针并递增重复指针
	di := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[di] = nums[i]
			di++
		}
	}
	return di
}

