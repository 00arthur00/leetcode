/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
func addBinary(a string, b string) string {

	result := make([]byte, 0, len(a)+len(b))
	var br byte
	for ia, ib := len(a)-1, len(b)-1; ia >= 0 || ib >= 0 || br > 0; ia, ib = ia-1, ib-1 {
		if ia >= 0 {
			br += a[ia] - '0'
		}
		if ib >= 0 {
			br += b[ib] - '0'
		}
		result = append(result, br%2+'0')
		br >>= 1
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
