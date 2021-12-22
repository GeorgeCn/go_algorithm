package answers

import "fmt"

/**
    题目 组合总和 中等
给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。

candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。 

对于给定的输入，保证和为 target 的唯一组合数少于 150 个

示例 1：

输入: candidates = [2,3,6,7], target = 7
输出: [[7],[2,2,3]]
*/

func main()  {
    ret := combinationSum([]int{2,3,6,7}, 7)
    fmt.Println("组合总和", ret)
}

func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}

    var dfs func(start int, temp []int, sum int)
    dfs = func(start int, temp []int, sum int) {
        if sum >= target {
            if sum == target {
                newTmp := make([]int, len(temp))
                copy(newTmp, temp)
                res = append(res, newTmp)
            }
            return
        }
        for i := start; i < len(candidates) ; i++  {
            temp = append(temp, candidates[i])
            dfs(i, temp, sum+candidates[i])
            temp = temp[:len(temp)-1]
            fmt.Println("temp:", temp,i)
        }
    }
    dfs(0, []int{}, 0)
    return res
}



