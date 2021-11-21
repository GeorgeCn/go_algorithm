package main

import "fmt"

func main()  {
    arr := []int{51,46,20,18,65,97,82,30,77,50,2,33,12,100}
    fmt.Println("冒泡排序前：", arr)

    selectSort(arr)

    fmt.Println("冒泡排序后：", arr)
}

func selectSort(arr []int) []int {
    for i := 0; i < len(arr)-1; i++  {
        min := i

        for j := i + 1; j < len(arr) ; j++  {
            if arr[min] > arr[j] {
                min = j
            }
        }
        // 每轮 最小的放到已排序列末尾
        arr[i], arr[min] = arr[min], arr[i]
    }

    return arr
}
