/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	ri := m + n - 1
	//双指针,从后往前每次赋值一个
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[ri] = nums1[i]
			i--
		} else {
			nums1[ri] = nums2[j]
			j--
		}
		ri--
	}
	//对j<0的情况不用处理,因为此时nums1中已经是正确值了
	if i < 0 {
		for ; j >= 0; j-- {
			nums1[ri] = nums2[j]
			ri--
		}
	}
}

