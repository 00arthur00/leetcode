/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */
func split(date string) (year, month, day int) {
	ss := strings.Split(date, "-")
	year, _ = strconv.Atoi(ss[0])
	month, _ = strconv.Atoi(ss[1])
	day, _ = strconv.Atoi(ss[2])
	return
}

func sepecailYear(year int) bool {
	//闰月判断标准:4的倍数且不是100的倍数,或者为400的倍数
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}

func dayOfYear(date string) int {
	year, month, day := split(date)
	days := 0
	//从2月份开始算,1月份的直接显示
	for i := 2; i <= month; i++ {
		switch i {
		case 2, 4, 6, 8, 9, 11:
			days += 31
		case 3:
			days += 28
		default:
			days += 30
		}
	}
	//2月份之后的才判断是否为闰月
	if month > 2 && sepecailYear(year) {
		return days + day + 1
	}
	return days + day
}
