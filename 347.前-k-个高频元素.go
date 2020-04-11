/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	valueCount := make(map[int]int)
	for _, n := range nums {
		valueCount[n] += 1
	}
	freqCount := make([]int, len(valueCount))
	for _, count := range valueCount {
		freqCount = append(freqCount, count)
	}
	topKFrequentByCount(freqCount, k)

}

func topKFrequentByCount(freqCount []int, left, right, k int) {
	if k <= 0 || right-left <= 1 || k > right-left {
		return
	}

	index := partition(freqCount, 0, len(freqCount)-1)
	if index == len(freqCount)-k {
		return
	} else if index < len(freqCount)-k {
		topKFrequent(freqCount, index+1, len(freqCount)-1, k)
	} else if index > len(freqCount)-k {
		topKFrequent(freqCount, 0, index-1, index+k-len(freqCount))
	}
}

func partition(nums []int, left, right int) int {
	key := nums[left]
	for left < right {
		for left < right && nums[right] >= nums[left] {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < nums[right] {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = key
	return left
}

// @lc code=end

