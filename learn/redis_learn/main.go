package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// 初始化redis
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	return err
}

// Set / Get操作
func redisDemo() {
	fmt.Printf("\n*****redisDemo*****\n")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	err1 := rdb.Set(ctx, "key", 100, 0).Err()
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	// 执行命令获取结果
	val, err := rdb.Get(ctx, "key").Result()
	fmt.Println(val, err) // 100 nil

	// 直接执行命令获取值
	value := rdb.Get(ctx, "key").Val()
	fmt.Println(value) // 100

	// 先获取到命令对象
	cmder := rdb.Get(ctx, "key")
	fmt.Println(cmder)       // get key:100
	fmt.Println(cmder.Val()) // 100
	fmt.Println(cmder.Err()) // nil
}

// HGetAll HMGet HGet
func hgetDemo() {
	fmt.Printf("\n*****hgetDemo*****\n")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	v1, err := rdb.HGetAll(ctx, "user").Result()
	if err != nil {
		fmt.Printf("hgetall failed, err:%v\n", err)
		return
	}
	fmt.Println(v1)

	v2 := rdb.HMGet(ctx, "user", "name", "age").Val()
	fmt.Println(v2)

	v3 := rdb.HGet(ctx, "user", "age").Val()
	fmt.Println(v3)
}

// zsetDemo 操作zset示例
func zsetDemo() {
	fmt.Printf("\n*****zsetDemo*****\n")
	zsetKey := "language_rank" // key
	languages := []*redis.Z{   // value
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// ZADD
	err := rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取分数最高的3个
	ret := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

// pipeDemo管道多条命令同时发送,节省了时间
func pipeDemo() {
	fmt.Printf("\n*****pipeDemo*****\n")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	pipe := rdb.Pipeline()

	incr := pipe.Incr(ctx, "pipeline_counter")      //给pipeline_counter这个key+1
	pipe.Expire(ctx, "pipeline_counter", time.Hour) //设置过期的时间是一个小时

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())
}

// watchDemo执行事务操作
func watchDemo() error {
	fmt.Printf("\n*****watchDemo*****\n")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel() //设定超时时间timeout后，返回的子Context的相关操作视为完成。

	key := "watch_count"
	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
	if err != nil {
		fmt.Println("exec failed...", err)
		return err
	}
	fmt.Println("exec successed!")
	return nil
}

func main() {
	if err := initClient(); err != nil {
		fmt.Println("redis start failed...", err)
	}
	fmt.Println("connect redis successed!")
	defer rdb.Close() //程序退出时释放相关资源
	redisDemo()
	hgetDemo()
	zsetDemo()
	pipeDemo()
	watchDemo()
}
