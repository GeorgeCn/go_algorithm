package answers

import "fmt"

/**
    题目 两数之和 简单

    给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

    示例

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

    进阶
写出时间复杂度小于o(n^2）的算法
 */
func main()  {
    ret := twoSum([]int{2,7,11,15}, 9)
    fmt.Println("twoSum:%T", ret)

    rett := twoSumSuper([]int{2,7,11,15}, 9)
    fmt.Println("twoSumSuper:%T", rett)
}


func twoSum(nums []int, target int) [][2]int  {
    //初始化二维数组
    arr := [][2]int{}

    for i, _ := range nums {
        for j := i + 1; j < len(nums); j++  {
            if target == nums[i] + nums[j] {
                arr = append(arr, [2]int{i, j})
            }
        }
    }

    return arr
}

// 进阶训练 使用map
func twoSumSuper(nums []int, target int) [][2]int {
    //初始化二维数组
    arr := [][2]int{}
    m := make(map[int]int)

    for index, val := range nums {
        if preindex, ok := m[target - val]; ok{
            arr = append(arr, [2]int{preindex, index})
        } else {
            m[val] = index
        }
    }

    return arr
}