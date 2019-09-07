/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
//1. 以i为结尾的连续子数组的最大和
//sum[i]=max{sum[i-1]+a[i],a[i]}
//sum[0]=nums[0]
//2. 从里面选出最大的来就是结果
func maxSubArray1(nums []int) int {
	sum := make([]int, len(nums), len(nums))
	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		sum[i] = max(nums[i]+sum[i-1], nums[i])
	}
	ans := sum[0]
	for i := 1; i < len(sum); i++ {
		ans = max(sum[0], ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSubArray(nums []int) int {
	sum := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i]+sum, nums[i])
		if sum > ans {
			ans = sum
		}
	}
	return ans
}


