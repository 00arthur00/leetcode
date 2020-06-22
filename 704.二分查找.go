/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	b,e:=0,len(nums)-1
	for b<=e{
		mid:=(b+e)/2
		if nums[mid]==target{
			return mid
		}else if nums[mid]>target{
			e=mid-1
		}else if nums[mid]<target{
			b=mid+1
		}
	}
	return -1
}
// @lc code=end

