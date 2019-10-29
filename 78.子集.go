/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
func subsets(nums []int) [][]int {
	arrayLen := len(nums)
	totalSetNum := 1 << uint(arrayLen)
	ans := make([][]int, 0, totalSetNum)
	ans = append(ans, []int{})
	for i := 0; i < len(nums); i++ {
		preTotal := len(ans)
		for j := 0; j < preTotal; j++ {
			ans = append(ans, append(ans[j], nums[i]))
		}
	}
	return ans
}

func subsets1(nums []int) [][]int {
	arrayLen := len(nums)
	totalSetNum := 1 << uint(arrayLen)
	ans := make([][]int, 0, totalSetNum)
	for i := 0; i < totalSetNum; i++ {
		result := make([]int, 0, arrayLen/2)
		for j := 0; j < arrayLen; j++ {
			if (1<<uint(j))&i > 0 {
				result = append(result, nums[j])
			}
		}
		ans = append(ans, result)
	}
	return ans
}

