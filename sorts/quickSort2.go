package main

import "fmt"

func main()  {
    arr := []int{51,46,20,18,65,97,82,30,77,50,2,33,12,100}
    fmt.Println("快速排序前：", arr)

    quickSort2(arr, 0, len(arr)-1)

    fmt.Println("快速排序后：", arr)

}

func quickSort2(arr []int, left, right int)  {
    if left < right {
        // 中间坐标
        mid := partion2(arr, left, right)

        quickSort2(arr, left, mid-1)
        quickSort2(arr, mid+1, right)
    }
}

// 分区过程 返回关键字落地索引
func partion2(arr []int, left, right int ) int {
    // 创建关键字
    pivot := arr[left]
    l := left
    r := right

    for {
        // 找到异常数字 退出循环
        for ;arr[l] <= pivot && l < r ;  {
            l++
        }

        for ;arr[r] > pivot ;  {
            r--
        }

        // 退出条件
        if l >= r {
            break
        }

        // 交换位置
        arr[l], arr[r] = arr[r], arr[l]
    }
    // 左边换到中间
    arr[l], arr[left] = arr[left], arr[l]

    return l
}
