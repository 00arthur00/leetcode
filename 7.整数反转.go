/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
func reverse(x int) int {
	const MAXINT = 1<<31 - 1
	const MININT = -1 << 31
	x64 := int64(x)
	var isNegative bool
	if x64 < 0 {
		isNegative = true
		x64 = -x64
	}

	var result int64
	for x64 > 0 {
		result = x64%10 + result*10
		x64 = x64 / 10
	}

	if isNegative {
		result = -result
	}
	if result < MININT || result > MAXINT {
		return 0
	}

	return int(result)
}

