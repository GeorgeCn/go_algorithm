package answers

import "fmt"

/**
    题目 二叉树的后序遍历 简单
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]

*/

func main()  {
    ret := postorderTraversal(&TreeNode{Val:1, Right:&TreeNode{Val:2, Left:&TreeNode{Val:3,Right:&TreeNode{Val:5}}}})
    fmt.Println("二叉树的后序遍历：", ret)
}

func postorderTraversal(root *TreeNode) []int {
    ret := []int{}
    var  order func(*TreeNode)
    order = func (node *TreeNode){
        if node ==nil{
            return
        }
        // 递归方法  先写入值 再递归左节点 再递归有节点
        ret = append(ret,node.Val)
        order(node.Left)
        order(node.Right)
    }
    order(root)
    return ret
}

// 递归算法 最小问题解
func dfs(node *TreeNode, temp []int) []int  {

    // 退出条件 异常
    if node == nil {
        return  temp
    }

    // 递归方法 先写入值 再递归左节点 再递归右节点
    temp = append(temp, node.Val)
    fmt.Println(temp)

    ret := dfs(node.Left, temp)
    rett := dfs(node.Right, temp)
    fmt.Println("二次",temp,ret,rett)

    return temp
}

 type TreeNode struct {
        Val int
        Left *TreeNode
        Right *TreeNode
 }


