package answers

import (
    "fmt"
    "sort"
)

/**
    题目 数组中的第K个最大元素 中等
    给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

 

示例:

输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/

func main() {
    ret := findKthLargest([]int{3,2,1,5,6,4}, 2)

    fmt.Println("数组中的第K个最大元素", ret)
}


func findKthLargest(nums []int, k int) int {
    sort.Sort(sort.Reverse(sort.IntSlice(nums)))

    return nums[k]
}


