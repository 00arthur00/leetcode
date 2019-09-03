/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	ans := len(prefix)
	for {
		var shouldMinus bool
		if ans == 0 {
			prefix = ""
			break
		}
		//以strs[0]为基准,从最长开始扫描
		prefix = prefix[0:ans]
		//对剩下的每个string查看是否相等,不相等的则直接退出
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < ans || strs[i][0:ans] != prefix {
				shouldMinus = true
				break
			}
		}
		//基准的长度减一,否则命中,break
		if shouldMinus {
			ans--
		} else {
			break
		}
	}
	return prefix
}

