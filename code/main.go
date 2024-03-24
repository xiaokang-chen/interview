package main

import (
	"fmt"
	"interview/code/list"
)

func main() {
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}
	list1 := list.ArrToList(nums1)
	list2 := list.ArrToList(nums2)

	resNode := list.AddTwoNumbers(list1, list2)
	res := list.ListToArr(resNode)
	fmt.Println("res", res)
}
