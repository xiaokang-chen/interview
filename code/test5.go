package main

import "math"

// GetMinAbsolute
// [-8, -5, -3, -1, 0, 2, 5]
func GetMinAbsolute(arr []int) int {
	left, right := 0, len(arr)-1
	ans := math.MaxInt
	isPositive := 1
	for left <= right {
		mid := left + (right-left)/2
		cur := arr[mid]
		if cur == 0 {
			return 0
		} else if cur > 0 {
			left = mid + 1
			isPositive = 1
			ans = min(ans, cur)
		} else if cur < 0 {
			right = mid - 1
			isPositive = -1
			ans = min(ans, -cur)
		}
	}
	return isPositive * ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Climb1
// i=3 a,b,c = 1,1,2 = f(1),f(2),f(3)
// i=4 a,b,c = 1,2,3 = f(2),f(3),f(4)
// i=5 a,b,c = 2,3,4 = f(3),f(4),f(5)
// i=6 a,b,c = 3,4,6 = f(4),f(5),f(6)
func Climb1(n int) int {
	a, b, c := 1, 1, 1
	for i := 2; i < n; i++ {
		a, b, c = b, c, a+c
	}
	return c
}

// 递归
// f(0) = 1
// f(1) = 1
// f(2) = 1
// f(3) = 2
// f(4) = 3
// f(5) = 4
// f(6) = 6
// a -> f(n-3)
// b -> f(n-1)
// c -> f(n)
// n=4 a,b,c = 1,1,3 = f(0),f(1),f(2)
// n=5 a,b,c = 1,3,4 = f(2),f(4),f(5)
// n=6 a,b,c = 2,4,6 = f(3),f(5),f(6)
// n=7 a,b,c = 3,6,9 = f(4),f(6),f(7)
func Climb2(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return 1
	}
	return Climb2(n-1) + Climb2(n-3)
}
