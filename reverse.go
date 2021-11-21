package main

import (
    "fmt"
    "math"
)

/**
    题目 整数反转 简单

    给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
 

示例 1：

输入：x = 123
输出：321

 */
func main() {

    ret := reverse(-1534236469)
    fmt.Println("reverse:", ret)
}


func reverse(x int) int  {

    if x < -1<<31 || x > (1<<31 -1) || x == 0{
        return 0
    }

    slice := []int{}
    var y int

    for {
        slice = append(slice, x%10)
        x = x/10

        if x == 0 {
            break
        }
    }

    length := len(slice)-1 // 索引长度-1
    for i, v := range slice {
        bite := int(math.Pow10(length - i))

        y += v * bite
    }

    if y < -1<<31 || y > (1<<31 -1) || y == 0{
        return 0
    }


    return y
}

