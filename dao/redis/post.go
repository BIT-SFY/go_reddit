package redis

import (
	"context"
	"reddit/models"

	"github.com/go-redis/redis/v8"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	ctx := context.Background()
	// 从redis获取id
	// 1.根据用户请求中携带的order参数来确定查询redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	// 2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	// 3.ZRevRange查询，按分数从大到小的顺序查询指定数量的元素
	return rdb.ZRevRange(ctx, key, start, end).Result()

}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	ctx := context.Background()
	// 使用pipeline一次执行多条命令，减少rtt
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		// 查找key中分数是1的元素的数量 -> 统计每篇帖子的赞成票的数量
		pipeline.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
