/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//两个指针:遍历指针和重复指针,相邻元素不相等的时候覆盖重复指针并递增重复指针
	di := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[di] = nums[i]
			di++
		}
	}
	return di
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//两个指针:遍历指针和重复指针,遍历指针与重复指针不相等的时候覆盖重复指针并递增重复指针
	di := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[di] {
			di++
			nums[di] = nums[i]
		}
	}
	return di + 1
}


