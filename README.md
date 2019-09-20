# leetcode
solution of golang
# backtracing
DFS+状态重置+剪枝
排列和组合的代码就差一行
``` golang
func helper(nums, trace []int, output *[][]int) {
	if len(nums) == 0 {
		ret := make([]int, len(trace))
		copy(ret, trace)
		*output = append(*output, ret)
		return
	}
	for i, n := range nums {
		trace = append(trace, n)
        nums[0], nums[i] = nums[i], nums[0]
        //排列
		helper(nums[1:len(nums)], trace, output)
        //组合
        helper(nums[i+1:len(nums)], trace, output)
        nums[0], nums[i] = nums[i], nums[0]
		trace = trace[0 : len(trace)-1]
	}
}
```