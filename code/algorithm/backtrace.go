package algorithm

// Permute 全排列
// 回溯算法就是纯暴力穷举，复杂度一般都很高
func Permute(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	used := make([]bool, len(nums))
	var backtrace func([]int, []int, []bool)
	// backtrace
	backtrace = func(nums []int, track []int, used []bool) {
		// 到达边界
		if len(nums) == len(track) {
			temp := make([]int, len(track))
			copy(temp, track)
			res = append(res, temp)
			return
		}
		for i := range nums {
			// 1.做选择
			if used[i] {
				continue
			}
			track = append(track, nums[i])
			used[i] = true
			// 2.回溯
			backtrace(nums, track, used)
			// 3.取消选择
			track = track[:len(track)-1]
			used[i] = false
		}
	}
	backtrace(nums, track, used)
	return res
}
