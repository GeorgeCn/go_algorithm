package answers

import "fmt"

/**
    题目 两数相加 简单

    给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]
*/

type ListNode struct {
    Val int
    Next *ListNode
}

func main()  {
    l1 := &ListNode{2,&ListNode{4, &ListNode{3, nil}}}
    l2 := &ListNode{5,&ListNode{6, &ListNode{4, nil}}}

    ret := addTwoNumbers(l1, l2)

    fmt.Printf("两数相加：%t", ret)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {

    // 定义位值和进位信号量
    var n1, n2 int
    var carry int
    var tail *ListNode

    for ;l1 != nil || l2 != nil; {
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }

        sum := n1 + n2 + carry
        sum, carry = sum%10, sum/10

        // tail实例化 将指针传递下去
        if head == nil {
            head = &ListNode{Val: sum}
            tail = head
        } else {
            tail.Next = &ListNode{Val: sum}
            tail = tail.Next
        }
    }

    if carry > 0 {
        tail.Next = &ListNode{Val: carry}
    }

    return
}