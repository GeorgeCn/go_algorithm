package answers

/**
    题目 旋转图像 中等
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

示例：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[[7,4,1],[8,5,2],[9,6,3]]

*/

func main()  {
    rotate([][]int{{1,2,3},{4,5,6},{7,8,9}})

}

func rotate(matrix [][]int) {
    n := len(matrix)
    tmp := make([][]int, n)
    for i := range tmp {
        tmp[i] = make([]int, n)
    }

    for i, row := range matrix {
        for j, v := range row {
            tmp[j][n-1-i] = v
        }
    }
    copy(matrix, tmp)
}


