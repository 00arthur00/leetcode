/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */
//自低向上dp解法
func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
func coinChange(coins []int, amount int) int{
	dp:=make([]int,amount+1)
	dp[0]=0
	for i:=1;i<=amount;i++{
		dp[i] = amount+1
	}
	for _,c:=range coins{
		for i:=c;i<=amount;i++{
			dp[i] = min(dp[i],dp[i-c]+1)
		}
	}
	if dp[amount]==amount+1{
		return -1
	}
	return dp[amount]
}

func coinChange2(coins []int, amount int) int{
	dp:=make([]int, amount+1)
	var i int 
	for i=1;i<=amount;i++{
		dp[i] = 1<<31
		for _,c:=range coins{
			if i-c < 0{
				continue
			}
			if dp[i-c]!=-1{
				dp[i] = min(1+dp[i-c],dp[i])
			}
		}
		if dp[i]==1<<31{
			dp[i] = -1
		}
	}
	return dp[amount]

}
//暴力解法
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func coinChangeRaw(coins []int, amount int) int{
	if amount ==0{
		return 0
	}
	if amount<0{
		return -1
	}
	var res int =  1<<31
	for _,c := range coins{
		subProm:=coinChange(coins,amount-c)
		if subProm ==-1{
			continue
		}
		res = min(res,1+subProm)
	}
	if res==1<<31{
		return -1
	}
	return res
}
//备忘录解法
func coinChangeMemo(coins []int, amount int) int{
	memo:=make(map[int]int)
	
	return dp(amount,memo,coins)
}
func dp(amount int,memo map[int]int, coins []int) int{
	
	if amount==0{
		return 0
	}
	if amount<0{
		return -1
	}

	if val,ok:=memo[amount];ok{
		return val
	}

	res:=1<<31
	
	for _,c := range coins{
		sub:=dp(amount-c,memo,coins)
		if sub==-1{
			continue
		}
		res = min(res,sub+1)
	}
	if res == 1<<31{
		res = -1
	}
	memo[amount]=res
	return res
}

/*
f(n)={
	0,  n=0
	min{f(n-ci)|{i in [1...k]}}, n>0
	-1, n<0
}
*/
