package redis

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis/v8"
)

// 基于用户投票的相关算法:http://www.ruanyifeng.com/blog/algorithm/
// 使用简化版的投票分数
// 用户投一票就加432分 86400/200 -> 需要两百张赞成票,就可以给帖子续一天 -> <<redis实战>>

/*投票的几种情况
1. direaction=1时,有两种情况:
	1.1 之前没有投过票,现在投赞成票	--> 更新分数和投票纪录 差值的绝对值: 1  +432
	1.2 之前投反对票,现在改投赞成票 --> 更新分数和投票纪录 差值的绝对值: 2  +432*2
2. direaction=0时,有两种情况:
	2.1 之前投过反对票,现在取消投票 --> 更新分数和投票纪录 差值的绝对值: 1  +432
	2.2 之前投过赞成票,现在取消投票 --> 更新分数和投票纪录 差值的绝对值: 1	-432
3. direaction=1时,有两种情况:
	3.1 之前没有投过票,现在投反对票 --> 更新分数和投票纪录 差值的绝对值: 1	-432
	3.2 之前投赞成票,现在改投反对票 --> 更新分数和投票纪录 差值的绝对值: 2  -432*2

投票的限制:
每个帖子自发表之日起,一个星期之内允许用户投票,超过一个星期就不允许用户投票了
	1. 到期之后将redis中保存的赞成票数以及反对票数存储到mysql中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600 // 一周有多少秒
	scorePerVote     = 432           // 一票有多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) (err error) {
	pipeline := rdb.TxPipeline()
	ctx := context.Background()
	// 帖子时间
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err = pipeline.Exec(ctx)
	return
}

func VoteForPost(userID, postID string, value float64) error {
	// 1.判断投票的限制
	// 去redis取帖子发布时间
	ctx := context.Background()
	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 2和3需要放到一个pipeline事务中操作
	// 2.更新帖子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值

	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	// 3.记录用户为该帖子投过票
	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postID), &redis.Z{
			Score:  value, //赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec(ctx)
	return err
}
