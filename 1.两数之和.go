/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, num := range nums {
		t2 := target - num
		if j, ok := lookup[t2]; ok {
			return []int{j, i}
		}
		lookup[num] = i
	}
	return []int{}
}

