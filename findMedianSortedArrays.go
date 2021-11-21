package main

import (
    "fmt"
)

/**
    题目 寻找两个正序数组的中位数 困难

        给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

 
示例：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
*/
func main()  {
    ret := findMedianSortedArrays([]int{1,3}, []int{2})
    fmt.Println("findMedianSortedArrays:", ret)
}


// 看时间复杂度要求知道 要用分治法解
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var slice []int
    m, n := devide(nums1, nums2)

    for _,v := range m {
        slice = append(slice, v)
    }

    for _,v := range n {
        slice = append(slice, v)
    }

    mid := (len(m) + len(n)) >> 1

    if len(slice) % 2 == 0 {
        return float64(slice[mid] + slice[mid+1]) / 2
    } else {
        return float64(slice[mid])

    }
}

func devide(nums1 []int, nums2 []int) ([]int, []int) {
    // 选取m, n的中位值
    m := len(nums1) >> 1
    n := len(nums2) >> 1

    if m == 0 || n == 0 {
        // 判断 大小
        return nums1, nums2

    } else {
        if nums1[0] <= nums2[0] {
            nums1 = nums1[:m]
            nums2 = nums2[0:n]
        } else {
            nums1 = nums1[0:n]
            nums2 = nums2[:m]
        }

        return devide(nums1, nums2)
    }

}
