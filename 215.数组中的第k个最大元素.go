/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
//在快排的过程中寻找第k大的数
func findKthLargest(nums []int, k int) int {
	index := helper(nums, 0, len(nums)-1, len(nums)-k)
	return nums[index]
}
func helper(ns []int, b, e, k int) int {
	pI := partition(ns, b, e)
	if pI < k {
		return helper(ns, pI+1, e, k)
	} else if pI > k {
		return helper(ns, b, pI-1, k)
	}
	return pI
}
func partition(ns []int, b, e int) int {
	key := ns[b]
	for b < e {
		for b < e && ns[e] > key {
			e--
		}
		ns[b] = ns[e]
		for b < e && ns[b] <= key {
			b++
		}
		ns[e] = ns[b]
	}
	ns[b] = key
	return b
}
