package main

import "fmt"

/**冒泡排序**/

func main()  {
    // 初始化数据
    arr := []int{51,46,20,18,65,97,82,30,77,50,2,33,12,100}
    fmt.Println("冒泡排序前：%T", arr)

    //bubbleSort1(arr)
    bubbleSort2(arr)

    fmt.Println("冒泡排序后：%T", arr)
}

// 每圈相邻元素对比
func bubbleSort1(arr []int) []int  {

    // 最后一个元素无需排序
    for i := 0; i < len(arr) -1; i++  {
        for j := i+1; j < len(arr) ; j++ {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }

    return arr
}

// 圈数迭代 每次从第一个元素开始
func bubbleSort2(arr []int) []int  {

    // 最后一个元素无需排序
    for i := 0; i < len(arr) -1; i++  {
        for j := 0; j < len(arr) -1 -i ; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    return arr
}

