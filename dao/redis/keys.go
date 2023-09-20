package redis

import "context"

// redis key注意使用命名空间的方式,方便查询和拆分

const (
	Prefix             = "reddit:"    // 项目key签注
	KeyPostTimeZSet    = "post:time"  // zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score" // zset;帖子及投票分数
	KeyPostVotedZSetPF = "post:voted" // zset;记录用户及投票类型;参数是:postid
	KeyCommunitySetPF  = "communit:"  // set;保存每个分区下帖子的id
)

var ctx = context.Background()

// getRedisKey 给key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
