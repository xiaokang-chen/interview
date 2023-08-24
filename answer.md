# answer

[toc]

## 一、逻辑题

1. D
2. B
3. C
4. C
5. A

## 二、语言特性

1. [0, 0, 1, 2]
2. A,B
3. C
4. C
5. B

## 三、计算机网络、操作系统、数据库

1. 3,4
2. cookie存储在客户端，session存储在服务端；cookie有大小限制（4kb），session没有限制；cookie通常用来存储用户信息，session用来存储登录状态；cookie和session配合一起解决Http请求无状态的问题；
3. 301：永久重定向；302：临时重定向；400：客户端请求出错；500：服务端响应出错；
4. grep：文本搜索；top：进程资源占用；find：目录下搜索文件；mv：文件移动；
cat：文件查看；wc：文件字数统计；chmod：文件权限修改；ps：查看进程信息；
5. rwx分别代表文件的读、写、执行
6. explain：sql语句查询解释（用来优化索引）；
7. 工厂（简单工厂、工厂方法、抽象工厂）模式；单例模式（懒汉、饿汉）；外观模式；适配器模式；观察者模式；

## 四、编程题

1. 

```go
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
```

2. 

```go
// Climb
// 0-0,1-1,2-1,3-2,4-3
// f(x) = f(x-1) + f(x-3)
// f(4) = f(3) + f(1)；f(3)+1 f(1)+3
func Climb(n int) int {
	a, b := 0, 1
	for i := 2; i < n; i++ {
		a, b = b, a + b
	}
	return b
}
```