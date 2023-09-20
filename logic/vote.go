package logic

import (
	"reddit/dao/redis"
	"reddit/models"
	"strconv"

	"go.uber.org/zap"
)

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
