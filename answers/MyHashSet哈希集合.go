package answers

import (
    "container/list"
    "fmt"
)

/**
    题目 计哈希集合 简单
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 MyHashSet 类：

void add(key) 向哈希集合中插入值 key 。
bool contains(key) 返回哈希集合中是否存在这个值 key 。
void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。

*/

func main()  {
    ret := Constructor()
    fmt.Println("计哈希集合", ret)
}

const base  = 225

type MyHashSet struct {
    data []list.List
}


func Constructor() MyHashSet {
    return MyHashSet{make([]list.List, base)}
}


func (this *MyHashSet) Add(key int)  {
    if !this.Contains(key) {
        h := this.hash(key)
        this.data[h].PushBack(key)
    }
}


func (this *MyHashSet) Remove(key int)  {
    h := this.hash(key)

    for e := this.data[h].Front(); e != nil ; e.Next()  {
        if e.Value.(int) == key {
            this.data[key].Remove(e)
        }
    }
}


func (this *MyHashSet) Contains(key int) bool {
    h := this.hash(key)

    for e := this.data[h].Front(); e!= nil; e.Next() {
        if e.Value.(int) == key {
            return true
        }
    }

    return false
}

// hash函数
func (* MyHashSet) hash(key int) int {
    return key % base

}




