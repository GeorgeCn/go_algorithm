package main

import "fmt"

func main()  {
    arr := []int{51,46,20,18,65,97,82,30,77,50,2,33,12,100}
    fmt.Println("快速排序前：", arr)

    quickSort(arr, 0, len(arr)-1)

    fmt.Println("快速排序后：", arr)
}

func quickSort(arr []int, left, right int)  {
    if left < right {
        // 寻找中间点坐标
        m := partition(arr, left, right)

        quickSort(arr, left, m-1)
        quickSort(arr, m+1, right)
    }
}

// 分区过程 返回关键字落地索引
func partition(arr []int, left, right int)  int {
    // 创建关键值
    pivot := arr[left]
    l := left
    r := right

    // 比pivot小的放左边，比pivot大大放右边
    for {
        // 从pivot 左边找到大于pivot的值  比pivot小等于就继续
        for ; arr[l] <= pivot && l < r; {
            l++
        }

        // 从pivot 右边找到小于pivot的值  比pivot大就继续
        for ; arr[r] > pivot; {
            r--
        }

        //l >= r 本次循环结束 break
        if l >= r {
            break
        }

        // 左右填坑
        arr[l], arr[r] = arr[r], arr[l]
    }
    arr[l], arr[left] = arr[left], arr[l]

    return l
}