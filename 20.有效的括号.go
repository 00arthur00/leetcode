/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	if len(s) == 0 {
		return true
	}
	stack := make([]byte, 0, len(s)+1/2)
	left2right := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	right2left := map[byte]byte{'}': '{', ']': '[', ')': '('}
	for i := 0; i < len(s); i++ {
		if _, ok := left2right[s[i]]; ok {
			stack = append(stack, s[i])
			continue
		}
		if left, ok := right2left[s[i]]; ok {
			if len(stack) == 0 {
				return false
			}
			if left != stack[len(stack)-1] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

