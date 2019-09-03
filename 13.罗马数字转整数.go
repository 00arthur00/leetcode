/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */
func romanToInt(s string) int {
	r2i := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	num := 0
	for i := 0; i < len(s); {
		//注意这里的 i<=len(s)-2
		if (s[i] == 'I' || s[i] == 'X' || s[i] == 'C') && i <= len(s)-2 {
			intdata, ok := r2i[s[i:i+2]]
			if ok {
				num += intdata
				i = i + 2
				continue
			}
		}
		num += r2i[s[i:i+1]]
		i = i + 1
	}
	return num
}

