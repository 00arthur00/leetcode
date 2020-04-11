/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	path := make([]int, 0, n)
	result := make([]byte, n, n)
	bt(n, 1, path, &k, result)
	return string(result)
}
func bt(n, start int, path []int, k *int, result []byte) bool {
	if start > n {
		*k = *k - 1
		if *k == 0 {
			for i := 0; i < len(path); i++ {
				result[i] = byte('0' + path[i])
			}
			return true
		}
		// for i, n := range path {
		// 	result[i] = byte('0' + n)
		// }
		return false
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		if bt(start+1, n, path, k, result) {
			return true
		}
		path = path[0 : len(path)-1]
	}
	return false
}

// @lc code=end

