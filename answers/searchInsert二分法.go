package answers

import "fmt"

/**
    题目 搜索插入位置 简单
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

示例 1:

输入: nums = [1,3,5,6], target = 5
输出: 2

*/

func main()  {
    ret := searchInsert([]int{1,3,5,6}, 0)
    fmt.Println("搜索插入位置", ret)
}


// 二分查找法
func searchInsert(nums []int, target int) int {
    l, r := 0, len(nums)-1
    ans := len(nums)
    for l <= r  {
        mid := (l+r) >> 1

        if target <= nums[mid] {
            ans = mid
            r = mid - 1
        }

        if target > nums[mid] {
            l = mid + 1
        }

    }

    return ans
}