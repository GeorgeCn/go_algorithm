package answers

import (
    "fmt"
    "sync"
)

/**
    题目 空模板 简单

*/
var wg sync.WaitGroup
var lock sync.Mutex
var strPool = sync.Pool{
    New: func() interface{}  {
        return "test str"
    },
}
func main() {
    str := strPool.Get()
    fmt.Println(str)
    strPool.Put("hello world")
    strr := strPool.Get()
    fmt.Println(strr)
    //m := sync.Map{}
    //wg.Add(20000)
    //go func() {				//开一个协程写map
    //    for i := 0; i < 10000; i++ {
    //        defer wg.Done()
    //        m.Store(i,i)
    //    }
    //}()
    //
    //go func() {				//开一个协程读map
    //    for i := 0; i < 10000; i++ {
    //        defer wg.Done()
    //        fmt.Println(m.Load(i))
    //    }
    //}()
    //wg.Wait()
    //time.Sleep(time.Second * 20)
}


//func blank(params []int) int {
//
//        ch := make(chan struct{}, 1)
//        go func() {
//            <-ch
//            // do something
//        }()
//        ch <- struct{}{}
//        // ...
//
//    //fmt.Println(len(ch))
//
//}


