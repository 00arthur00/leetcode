/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	minPrice:= math.MaxInt64
	max:=0
	for _,price:=range prices{
		if minPrice>price{
			minPrice = price
			continue
		}
		if price-minPrice>max{
			max = price-minPrice
		}
	}
	return max
}
// @lc code=end

