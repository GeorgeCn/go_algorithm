package main

import "fmt"

/**
    题目 移除元素 简单

    实现 strStr() 函数。

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。

 

说明：

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。

 

示例 ：

输入：haystack = "hello", needle = "ll"
输出：2
 */
func main()  {
    ret := strStr("hello", "ll")
    fmt.Println("strStr", ret)
}

func strStr(haystack string, needle string) int {
    // 过滤边界问题
    if needle == "" {
        return 0
    }

    if len(needle) > len(haystack) {
        return -1
    }

    var l, r int

    for i := 0; i < len(haystack) - len(needle) +1; i++ {

        // 校验字串
        l = i
        r = i+len(needle)

        if r > len(haystack) {
            if haystack[l:] == needle {
                return i
            }
        } else {
            if haystack[l:r] == needle {
                return i
            }
        }
    }

    return -1
}