/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	//cache sum%n to index
	hash:=make(map[int]int)
	hash[0]=-1

	sum:=0
	for i,num:=range nums{
		sum += num
		if k!=0{
			sum = sum % k
		}

		bingle,ok:=hash[sum]
		if ok{
			if i-bingle > 1{
				return true
			}
			continue
		}
		
		hash[sum] = i
	}
	return false
}
// @lc code=end

