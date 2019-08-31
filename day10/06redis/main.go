package main

import (
	"fmt"
	"github.com/go-redis/redis"

)

var redisdb *redis.Client

func initRedis()(err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	_,err = redisdb.Ping().Result()
	return

}

func zset() {
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90,Member: "php"},
		&redis.Z{Score: 95,Member: "golang"},
		&redis.Z{Score: 97,Member: "java"},
		&redis.Z{Score: 99,Member: "python"},

	}
	//把元素追加到key
	redisdb.ZAdd(key,items...)
}

func zsetjia(){
	key := "rank"
	newScore ,err := redisdb.ZIncrBy(key,10,"python").Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("python's score is %f now \n",newScore)
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("链接redis成功")
	//zset()
	zsetjia()

}
