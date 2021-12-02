package main

import (
    "fmt"
    editor "leetcode/editor/cn"
)

func main()  {
    ret := editor.MajorityElement([]int{2,2,1,1,1,2,2})

    fmt.Println("多数元素", ret)
}
