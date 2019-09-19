/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */
func reverseBits(num uint32) uint32 {
	var res uint32
	for i, j := uint32(0), uint32(31); i < j; i, j = i+1, j-1 {
		left := uint32(1 << i)
		right := uint32(1 << j)
		if left&num == left {
			res |= right
		}
		if right&num == right {
			res |= left
		}
	}
	return res
}

