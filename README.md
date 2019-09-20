# leetcode
solution of golang
# backtracing
DFS+状态重置+剪枝
结果可通过树状结构遍历得到,如排列组合问题(无剪枝操作),组合求和问题
排列和组合的代码就差一行
``` golang
func helper(nums, trace []int, output *[][]int) {
    //退出条件
	if len(nums) == 0 {
		ret := make([]int, len(trace))
		copy(ret, trace)
		*output = append(*output, ret)
		return
	}
	for i, n := range nums {
        //状态
		trace = append(trace, n)
        nums[0], nums[i] = nums[i], nums[0]
        //排列
		helper(nums[1:len(nums)], trace, output)
        //组合
        helper(nums[i+1:len(nums)], trace, output)
        //重置状态
        nums[0], nums[i] = nums[i], nums[0]
		trace = trace[0 : len(trace)-1]
	}
}
```