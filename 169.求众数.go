/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 */
//1.排序的同时进行查找
func majorityElement1(nums []int) int {
	helper(nums, 0, len(nums)-1)
	return nums[len(nums)/2]
}
func helper(nums []int, l, r int) {
	if l >= r {
		return
	}
	pIndex := partition(nums, l, r)
	if pIndex < len(nums)/2 {
		helper(nums, pIndex+1, r)
	} else if pIndex > len(nums)/2 {
		helper(nums, l, pIndex-1)
	}
}
func partition(nums []int, l, r int) int {
	key := nums[l]
	for l < r {
		for l < r && nums[r] > key {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] <= key {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = key
	return l
}

//2. Boyer-Moore 投票算法:如果我们把众数记为 +1+1 ，把其他数记为 -1−1 ，将它们全部加起来，显然和大于 0 ，从结果本身我们可以看出众数比其他数多
func majorityElement(nums []int) int {
	count, candidate := 0, nums[0]
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count += 1
		} else {
			count -= 1
		}
	}
	return candidate
}
