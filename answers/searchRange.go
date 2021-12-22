package answers

import (
    "fmt"
    "sort"
)

/**
    题目 在排序数组中查找元素的第一个和最后一个位置 中等
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：

你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]


*/

func main()  {
    ret := searchRange([]int{5,7,7,8,8,10}, 8)

    fmt.Printf("在排序数组中查找元素的第一个和最后一个位置：%t", ret)
}

func searchRange(nums []int, target int) []int {
    m := make(map[int]int)
    left, right := -1, -1
    counter := 0

    for i, v := range nums{
        if v == target {
            fmt.Println(123, counter)

            if _, ok := m[v]; ok {
                if counter > 0 {
                    right = i
                }
            } else {
                m[v] = i
                left = i
                right = i
            }
            counter++
        }
    }

    return []int{left, right}
}

// 二分法
func searchRangeSuper(nums []int, target int) []int {
    leftmost := sort.SearchInts(nums, target)
    if leftmost == len(nums) || nums[leftmost] != target {
        return []int{-1, -1}
    }
    rightmost := sort.SearchInts(nums, target + 1) - 1
    return []int{leftmost, rightmost}
}