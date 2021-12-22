package answers

import (
    "context"
    "github.com/go-redis/redis/v8"
    "github.com/taoshihan1991/miaosha/setting"
    "log"
    "sync"
    "time"
)

var rdb *redis.Client
var ctx = context.Background()
var mutex sync.Mutex

func NewRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     setting.Redis.Ip + ":" + setting.Redis.Port,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}
func Lock(key string) bool {
    mutex.Lock()
    defer mutex.Unlock()
    bool, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
    if err != nil {
        log.Println(err.Error())
    }
    return bool
}
func UnLock(key string) int64 {
    nums, err := rdb.Del(ctx, key).Result()
    if err != nil {
        log.Println(err.Error())
        return 0
    }
    return nums
}