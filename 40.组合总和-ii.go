import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */
func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0)
	sort.Ints(candidates)
	bt(comb, candidates, target, &ans)
	return ans

}
func bt(comb, cans []int, target int, output *[][]int) {
	if target == 0 {
		ret := make([]int, len(comb))
		copy(ret, comb)
		*output = append(*output, ret)
	} else if target < 0 {
		return
	}
	for i, n := range cans {
		if i > 0 && cans[i] == cans[i-1] {
			continue
		}
		comb = append(comb, n)
		bt(comb, cans[i+1:len(cans)], target-n, output)
		comb = comb[0 : len(comb)-1]
	}
}
