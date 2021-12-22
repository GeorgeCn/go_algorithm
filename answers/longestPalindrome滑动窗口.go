package answers

import "fmt"

/**
    题目 最长回文子串 回文串（palindromic string）是指这个字符串无论从左读还是从右读，所读的顺序是一样的 困难

    给你一个字符串 s，找到 s 中最长的回文子串。

    示例 ：

    输入：s = "babad"
    输出："bab"
    解释："aba" 同样是符合题意的答案。

*/
func main() {

    ret := longestPalindrome("aaaaaaaa")
    fmt.Println("longestPalindromeSuper:", ret)
}


// 暴力破解 双循环方式 省略


// 使用map 做窗口平移
func longestPalindrome(s string) string {
    m := make(map[string]int)
    var temp string
    var length, maxLength int

    for index, v := range s {

       if preindex, ok := m[string(v)]; ok {
           length = index - preindex

           if length > maxLength {
               // 判断字串是否是回文字串

               pal := palindrome(s[preindex:index+1])

               if len(pal) < 3 {
                   if  len(pal) == 2{
                       if pal[0] != pal[1] {
                           continue
                       }
                   }
                   maxLength = length
                   temp = s[preindex:index+1]

               }
           }
       } else {
           m[string(v)] = index
       }
    }
    //处理边界问题
    if maxLength == 0 {
       temp = s[0:1]
    }

    return temp
}

func palindrome (s string) string {
    n := len(s)

    if n <= 2 {
        return s
    }

    // 去除首尾
    temp := s[1:n-1]

    if len(temp) != 1 && len(temp) > 0 {
        l := len(temp) -1
        if temp[0] == temp[l] {
            palindrome(temp)
        }
    }

    return temp
}

func max(x, y int) int {
    if x < y {
        return y
    }

    return x
}