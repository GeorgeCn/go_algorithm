package answers

import (
    "fmt"
    "math"
)

/**
    题目 回文数 简单

    给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

 

示例：

输入：x = 121
输出：true

 */
func main() {

    ret := isPalindrome(-111)
    fmt.Println("isPalindrome:", ret)
}


func isPalindrome(x int) bool  {

    if x >= 0 && x == reversePalindr(x) {
        return true
    }

    return false
}

// 整数反转
func reversePalindr(x int) int  {

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

