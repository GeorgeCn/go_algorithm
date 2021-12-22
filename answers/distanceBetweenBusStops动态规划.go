package answers

import (
    "fmt"
)

/**
    题目 公交站间的距离 简单
环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。

环线上的公交车都可以按顺时针和逆时针的方向行驶。

返回乘客从出发点 start 到目的地 destination 之间的最短距离。

*/

func main()  {
    ret := distanceBetweenBusStops([]int{7,10,1,12,11,14,5,0}, 7, 2)
    fmt.Println("公交站间的距离", ret)
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
    // 处理异常情况
    if start == destination {
        return 0
    }
    length := len(distance)
    dp := make([][]int, length)
    sumPath := 0

    for i := 0; i< length ;i++  {
        temp := make([]int, length)
        for j := 0; j < length;j++  {
            if j == i {
                temp[j] = 0
            }
        }
        sumPath += distance[i]
        dp[i] = temp
    }

    // 动态规划
    // 先做步长枚举

    for D:= 1; D < length; D++ {
        // 做距离坐标迭代
        for i:=0; i < length;i++ {
            // 根据步长换算j
            j := D + i

            // 边界条件
            if j >= length{
                break
            }

            // 状态转换
            if j- i < 2 {
                dp[i][j] = distance[i]
            } else {
                dp[i][j] = dp[i][i+1] + dp[i+1][j]
            }

        }

    }

    // 最短距离技术
    if start > destination {
        start, destination = destination, start
    }
    return mmin(dp[start][destination], sumPath - dp[start][destination])
}

func mmin(x, y int) int {
    if x < y {
        return x
    }

    return y
}


