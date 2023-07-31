# Golang

[toc]

## 一、golang 中 make 和 new 的区别

1）作用变量类型不同，new给string,int和数组分配内存，make给切片，map，channel分配内存；
2）返回类型不一样，new返回指向变量的指针，make返回变量本身；
3）new 分配的空间被清零。make 分配空间后，会进行初始化；

## 二、数组和切片的区别

1）定义方式不一样；
2）初始化方式不一样，数组需要指定大小，大小不改变；
3）在函数传递中，数组切片都是值传递；

数组和切片的定义及初始化：

```golang
// 数组
var a1 [3]int
var a2 [...]int{1,2,3}
a3 := [5]int{1,2,3}
var a4 = new([4]int)

// 切片
var a1 []int
var a2 := make([]int, 3, 5) // len=3; cap=5
```

## 三、rune类型是什么，怎么用

rune是int32类型的别名，在各方面都等价于它，用来区分字符串和整数值，使用单引号定义。用途如下：

1. 统计带中文字符串的长度

    ```go
    fmt.Println(len("Go编程语言"))         // 14
    fmt.Println(len([]rune("Go编程语言"))) // 6
    ```

2. 截取带中文的字符串

    ```go
    // 对于包含中文的字符串，如果用rune，可以很方便的根据
    // 需要截取的位置进行对包含中文字符串进行截断
    s := "Go编程语言"
    fmt.Println(s[:8])                 // Go编程
    fmt.Println(s[:7])                 // Go编�
    fmt.Println(string([]rune(s)[:4])) // Go编程
    ```

## 四、golang 中解析 tag 是怎么实现的？反射原理是什么？

Go 中解析的 tag 是通过反射实现的，反射将接口变量转换成反射对象 Type 和 Value。
反射可以通过反射对象 Value 还原成原先的接口变量；反射可以用来修改一个变量的值。
tag是啥:结构体支持标记。
