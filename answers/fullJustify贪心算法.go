package answers

import (
    "fmt"
    "strings"
)

/**
    题目 文本左右对齐 困难
给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

示例:

输入:
words = ["This", "is", "an", "example", "of", "text", "justification."]
maxWidth = 16
输出:
[
   "This    is    an",
   "example  of text",
   "justification.  "
]

*/

func main()  {
    ret := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
    fmt.Println("空模板", ret)
}

func fullJustify(words []string, maxWidth int) []string {
    ans := []string{}
    right, n := 0, len(words)

    for {
        left := right // 当前行的第一个单纯在words的位置

        sumLen := 0 //统计这行单词数量
        //循环确定当前行放多少个
        for right < n && sumLen+len(words[right])+ right-left <= maxWidth {
            sumLen += len(words[right])
            right++
        }

        // 当前行是最后一个：单词左对齐，单词间只有一个空，剩余填充空格
        if right == n {
            s := strings.Join(words[left:], " ")
            ans = append(ans, s +bblank(maxWidth- len(s)))
        }

        numWords := right -left
        numSpace := maxWidth - sumLen

        // 当前行只有一个单词
        if numWords == 1 {
            ans = append(ans, words[left] + bblank(numSpace))
            continue
        }

        // 当前不止一个单词
        avgSpace := numSpace /( numWords -1)
        extraSpace := numSpace % (numWords -1)

        // 拼接额外加一个空格的单词
        s1 :=strings.Join(words[left:left+extraSpace+1], bblank(avgSpace+1))
        s2 := strings.Join(words[left+extraSpace+1:right], bblank(avgSpace))  // 拼接其余单词

        ans = append(ans, s1+bblank(avgSpace)+s2)
    }
}

func bblank(n int) string {
    s := ""
    for n >= 0 {
        s += " "
        n--
    }

    return s
}


