package main

import "fmt"

func main() {

    arr := []int{51, 46, 20, 18, 65, 97, 82, 30, 77, 50, 2, 33, 12, 100}
    fmt.Println("归并排序前：", arr)

    mergeSort(arr, 0, len(arr)-1)

    fmt.Println("归并排序后：", arr)
}

func merge(arr []int, left, right, mid int) {
    var temp []int // 辅助数组
    // temp := []int{}
    p1 := left
    p2 := mid + 1

    for p1 <= mid && p2 <= right {
        if arr[p1] <= arr[p2] {
            temp = append(temp, arr[p1])
            p1++
        } else {
            temp = append(temp, arr[p2])
            p2++
        }
    }

    for p1 <= mid {
        temp = append(temp, arr[p1])
        p1++
    }

    for p2 <= right {
        temp = append(temp, arr[p2])
        p2++
    }

    // 将排序后的值拷贝到原数组
    for i, val := range temp {
        arr[left+i] = val
    }
}

func mergeSort(arr []int, left, right int) {
    if left < right {
        mid := (left + right) >> 1
        mergeSort(arr, left, mid)
        mergeSort(arr, mid+1, right)
        merge(arr, left, right, mid)
    }
}