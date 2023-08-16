package main

import (
	"fmt"
	"interview/code/list"
<<<<<<< Updated upstream
	"runtime"
=======
	"sort"
	"strconv"
>>>>>>> Stashed changes
	"sync"
)

// 线程池
type GoPool struct {
	wg  sync.WaitGroup
	arr []chan byte
}

// 交替打印26个字母
func PrintLetter() {
	count := 26
	wg := sync.WaitGroup{}
	// 切片，len为100
	gp := &GoPool{wg, make([]chan byte, 100)}

	for i := 0; i < count; i++ {
		gp.arr[i] = make(chan byte)
		// 最后一个chan初始化，作为启动，等待被读
		if i == count-1 {
			go func(i int) {
				gp.arr[i] <- 'A'
			}(i)
		}
	}
	fmt.Printf("协程数量：%d\n", runtime.NumGoroutine())
	gp.wg.Add(count)
	for i := 0; i < count; i++ {
		lastChan := make(chan byte)
		curChan := make(chan byte)
		if i == 0 {
			lastChan = gp.arr[count-1]
		} else {
			lastChan = gp.arr[i-1]
		}
		go func(lastChan, curChan chan byte) {
			defer gp.wg.Done()
			for i := 0; i < 5; i++ {
				s := <-lastChan
				fmt.Printf("%c\n", s)
				curChan <- 'A' + byte(i)
			}
		}(lastChan, curChan)
		fmt.Printf("%d, 协程数量：%d\n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("结束, 协程数量：%d\n", runtime.NumGoroutine())
}

// 链表判断环
// 快慢指针
func IsCircle(node *list.ListNode) bool {
	slow, fast := node, node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 快慢指针相遇，存在环
		if fast == slow {
			return true
		}
	}
	return false
}

// 三数之和
// 假设：不存在不相同的
// 定1找2
func ThreeSum(nums []int) [][]int {
	// cap, len
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		k := len(nums) - 1
		target := -1 * nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			// 向左移动k（减小nums[k]的大小）
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			// 指针重合，退出循环。代表没有在i固定情况下，满足条件的j
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}

// CountAndSay 38.外观数
func CountAndSay(n int) string {
	if n < 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	str := CountAndSay(n - 1)
	byteArr := []byte(str)
	ans := ""
	t := 0
	for i := 0; i < len(str); i++ {
		item := str[i]   // 49 = 0x31
		ru := rune(item) // 49
		i32 := int(item) // 49
		s := fmt.Sprintf("%c", item)
		char := strconv.Itoa(int(item))
		_ = item
		_ = ru
		_ = i32
		_ = s
		_ = char
	}
	for i := 0; i < len(byteArr); i += t {
		item := byteArr[i]
		t = 0
		for j := i; j < len(byteArr); j++ {
			if item == byteArr[j] {
				t++
			} else {
				break
			}
		}
		s := fmt.Sprintf("%d%c", t, item)
		ans += s
	}
	return ans
}
