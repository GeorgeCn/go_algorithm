package answers

import (
    "fmt"
)

/**
    题目 盛最多水的容器 中等

    给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。



 */
func main() {

    ret := maxArea([]int{4,3,2,1,4})
    fmt.Println("maxArea:", ret)
}



// 双指针移动
func maxArea(height []int) int  {
    l := 0
    r := len(height) -1
    maxArea := 0

    for(l < r) {
        area := area(height[l], height[r], l, r)
        maxArea = max(maxArea, area)

        if height[l] <= height[r] {
                l++
        } else {
                r--
        }
    }

    return maxArea
}

func min(x, y int) int  {
    if x > y {
        return y
    }

    return x
}

func max(x, y int) int  {
    if x < y {
        return y
    }

    return x
}

func area(x, y, l, r int) int {
    heigh := min(x, y)
    length := r - l

    area := int(heigh) * length

    return area
}

