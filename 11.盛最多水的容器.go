/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
func maxArea(height []int) int {
	b, e := 0, len(height)-1
	ans := 0
	for b < e {
		area := min(height[b], height[e]) * (e - b)
		if area > ans {
			ans = area
		}
		if height[b] < height[e] {
			b++
		} else {
			e--
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
