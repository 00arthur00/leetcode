/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0)
	backtracing(comb, candidates, target, &ans)
	return ans
}
func backtracing(comb []int, ns []int, target int, output *[][]int) {
	if target == 0 {
		result := make([]int, len(comb))
		copy(result, comb)
		*output = append(*output, result)
		return
	} else if target < 0 {
		return
	}
	for i, n := range ns {
		ttmp := target
		combLen := len(comb)
		for ttmp > 0 {
			ttmp -= n
			comb = append(comb, n)
			backtracing(comb, ns[i+1:len(ns)], ttmp, output)
		}
		comb = comb[0:combLen]
	}
}
