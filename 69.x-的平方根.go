/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */
func mySqrt(x int) int {
	b := 1
	e := x
	var mid int
	for {
		mid = (b + e) / 2
		if b == mid {
			break
		}
		res := mid * mid
		if res == x {
			break
		} else if res > x {
			e = mid
		} else {
			b = mid
		}
	}
	return mid
}

