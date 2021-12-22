package answers

import (
    "fmt"
)

/**
    题目 无重复字符的最长子串 中等

    给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

    示例 

    输入: s = "abcabcbb"
    输出: 3
    解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */
func main()  {

    ret := lengthOfLongestSubstring("abcdabcbb")
    fmt.Println("lengthOfLongestSubstring:", ret)

    rett := lengthOfLongestSubstringSuper("abcdabcbb")
    fmt.Println("lengthOfLongestSubstring:", rett)
}

//暴力解 时间复杂度O(n^2)
func lengthOfLongestSubstring(s string) int {

    maxCounter := 0
    for i := 0; i < len(s)-1; i++ {
        counter := 0
        for j := i+1; j< len(s); j++ {
            a := string(s[i])
            b := string(s[j])

            counter++
            // max函数实现
            if maxCounter < counter {
                maxCounter, counter = counter,maxCounter
            }

            if a == b {

                break
            }
        }
    }

    return  maxCounter
}

// 标签：滑动窗口
// 暴力解法时间复杂度较高，会达到 O(n^2)
// 故而采取滑动窗口的方法降低时间复杂度
// 定义一个 map 数据结构存储 (k, v)，其中 key 值为字符，value 值为字符位置 +1，加 1 表示从字符位置后一个才开始不重复
// 我们定义不重复子串的开始位置为 start，结束位置为 end
// 随着 end 不断遍历向后，会遇到与 [start, end] 区间内字符相同的情况，此时将字符作为 key 值，获取其 value 值，并更新 start，此时 [start, end] 区间内不存在重复字符
// 无论是否更新 start，都会更新其 map 数据结构和结果 ans。
// 时间复杂度：O(n)
func lengthOfLongestSubstringSuper(s string) int {
    m := make(map[string]int)
    maxLength := 0

    for i, v := range s{

        if start, ok := m[string(v)]; ok {
            m[string(v)] = i
            length := i - start +1
            if maxLength < length {
                maxLength, length = length, maxLength
            }
        } else {
            m[string(v)] = i // map值可以被覆盖
        }
    }

    if maxLength == 0{
        maxLength = len(s)
    }

    return  maxLength
}

// 用哈希表保存当前滑动窗口窗口的元素（保存每个元素索引）， 如果窗口的下一个元素已经在哈希表中，则将窗口的起点重置为哈希表中重复元素的下一个索引（跳过重复元素）。每一次迭代都通过max函数比较选出滑动窗口的最大长度
