package main

func majorityElement(nums []int) int {
	ans := nums[0]
	count := 0
	for _, item := range nums {
		// 如果当前遍历元素等于最终结果，则count++
		// 否则，count--
		if item == ans {
			count++
		} else {
			count--
		}
		// 如果当前元素和结果不一致并且
		if item != ans && count == 0 {
			ans = item
		}
	}
	return ans
}
