# Golang

[toc]

## 一、rune类型是什么，怎么用

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
