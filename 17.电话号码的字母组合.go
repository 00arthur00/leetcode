/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
func letterCombinations(digits string) []string {

	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	ans = append(ans, "")
	d2l := []string{
		"",     //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	for i := 0; i < len(digits); i++ {
		d := digits[i]
		ansLen := len(ans)
		for j := 0; j < ansLen; j++ {
			for _, l := range d2l[d-'0'] {
				ans = append(ans, ans[j]+string(l))
			}
		}
		ans = ans[ansLen:len(ans)]
	}
	return ans
}

//backtracing
func letterCombinations1(digits string) []string {
	d2l := []string{
		" ",    //0
		"",     //1
		"abc",  //2
		"def",  //3
		"ghi",  //4
		"jkl",  //5
		"mno",  //6
		"pqrs", //7
		"tuv",  //8
		"wxyz", //9
	}
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	backtrace("", digits, 0, &ans, d2l)
	return ans
}
func backtrace(comb string, digits string, index int, ans *[]string, d2l []string) {
	if index == len(digits) {
		*ans = append(*ans, comb)
		return
	}
	d := digits[index]
	for _, b := range d2l[d-'0'] {
		backtrace(comb+string(b), digits, index+1, ans, d2l)
	}
}
