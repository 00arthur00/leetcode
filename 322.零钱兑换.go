/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	ans := make([]int, amount+1, amount+1)
	//设置最大值为amount+1,最少硬币数必然小于amount+1,因为面额为int类型,最小值为1
	for i := 1; i <= amount; i++ {
		ans[i] = amount + 1
	}
	//自底向上,从1开始计算 f(n)
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			//币值比amount小,没有遍历必要
			if coin <= i {
				//取最小值并更新
				if ans[i-coin]+1 < ans[i] {
					ans[i] = ans[i-coin] + 1
				}
			}
		}
	}
	//如果为amount+1,则无解,return -1
	if ans[len(ans)-1] > amount {
		return -1
	}
	//return结果
	return ans[len(ans)-1]
}

/*
f(n)={
	0,  n=0
	min{f(n-ci)|{i in [1...k]}}, n>0
	-1, n<0
}
*/
