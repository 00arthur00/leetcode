/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
//https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode/
//从后往前找到第一个peek元素，然后反转，peek之后的元素，然后找到第一个比peek之前的元素大的元素，进行swap
func nextPermutation(nums []int) {
	index := findLastBackwardPeek(nums)
	reverse(nums[index:len(nums)])
	if index != 0 {
		ans := index - 1
		for i := index; i < len(nums); i++ {
			if nums[ans] < nums[i] {
				nums[ans], nums[i] = nums[i], nums[ans]
				break
			}
		}
	}
}
func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
func findLastBackwardPeek(nums []int) int {
	ans := len(nums) - 1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			ans = i - 1
		} else {
			break
		}
	}
	return ans
}

//从前往后找到最后一个升序序列，此元素之后的必为降序序列，反转降序序列的所有元素反转，
//然后找到降序序列之前的位置替换反转之后的降序序列里的第一个比他大的元素
//此方法比从后往前找第一个降序慢
func nextPermutation0(nums []int) {
	index := findLastAsc0(nums)
	if index == -1 {
		for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	} else {
		for i := index; i < len(nums); i++ {
			if nums[i] > nums[index] {
				nums[i], nums[index] = nums[index], nums[i]
				break
			}
		}
	}
}
func findLastAsc0(nums []int) int {
	lastAsc := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			lastAsc = i
		}
	}
	if lastAsc == -1 {
		return lastAsc
	}
	for i, j := lastAsc+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return lastAsc
}
