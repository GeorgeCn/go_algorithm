package answers

import (
    "container/list"
    "fmt"
)

/**
  题目 设计哈希映射 简单
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

实现 MyHashMap 类：

MyHashMap() 用空映射初始化对象
void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。

*/

func main()  {
    ret := blank()
    fmt.Println("空模板", ret)
}

const base  = 255
type entry struct {
    key, value int
}

type MyHashMap struct {
    data []list.List // 双向链表组
}


func ConstructorMap() MyHashMap {
    return MyHashMap{make([]list.List, base)}
}

// 更新
func (this *MyHashMap) Put(key int, value int)  {
    if this.Contains(key) {
       // 更新
        this.Remove(key)
        this.Add(key,value)
    } else {
        // 新增
        this.Add(key,value)
    }
}


func (this *MyHashMap) Get(key int) int {
    if this.Contains(key) {
        h := this.hash(key)

        for e := this.data[h].Front(); e != nil ; e = e.Next()  {
            if  et := e.Value.(entry); et.key == key{
                return e.Value.(entry).value
            }
        }
    }

    return -1
}

func (this *MyHashMap) Add(key, val int) {
    if !this.Contains(key) {
        h := this.hash(key)
        this.data[h].PushBack(entry{key, val})
    }
}


func (this *MyHashMap) Remove(key int)  {

    if this.Contains(key) {
        h := this.hash(key)
        for e := this.data[h].Front(); e != nil ; e.Next()  {
            if e.Value.(entry).value == key {
                this.data[h].Remove(e)
            }
        }
    }

}

func (this *MyHashMap) Contains(key int) bool {
    h := this.hash(key)
    // 遍历
    for e := this.data[h].Front(); e != nil ; e.Next()  {
        if e.Value.(entry).value == key {
            return true
        }
    }

    return false
}

func (this *MyHashMap) hash(key int) int {
    return key % base
}
