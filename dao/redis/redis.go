/*
 * @Author: shenfuyuan
 * @Date: 2023-09-14 17:25:02
 * @LastEditTime: 2023-09-16 08:54:40
 * @Description:
 */
package redis

import (
	"fmt"
	"reddit/settings"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

// 初始化redis
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() {
	rdb.Close()
}
