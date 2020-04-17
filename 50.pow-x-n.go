/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */
func myPow(x float64, n int) float64 {
	if n==0{
		return 1.0
	}
	ans, m := 1.0, n
	if n < 0 {
		m = -n
	}

	factor := x
	for m/2 > 0 {
		if m&1 == 1 {
			ans *= factor
		}
		factor *= factor
		m = m / 2
	}
	ans *=factor
	if n < 0 {
		return 1.0/ ans
	}
	return ans
}

