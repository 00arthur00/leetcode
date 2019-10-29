/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0, 0)
	ans = append(ans, []int{})
	//计算上轮开始的地方
	ansLoc := 0
	for i := 0; i < len(nums); i++ {
		b := 0
		if i > 0 && nums[i] == nums[i-1] {
			//重复数据从上轮开始的地方开始
			b = ansLoc
		}
		lenAns := len(ans)
		for ; b < lenAns; b++ {
			ans = append(ans, append(ans[b], nums[i]))
		}
		ansLoc = lenAns
	}
	return ans
}

