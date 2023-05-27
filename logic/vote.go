package logic

import (
	"go.uber.org/zap"
	//"github.com/go-redis/redis/v8"
	"liaoBa/dao/redis"
	"liaoBa/models"
	"strconv"
)

// 投票功能:
// 1. 用户投票的数据
// 2.

// 投一票+432分 86400/200 -》 需要200张赞成票，可以续一天
/* 投票情况分析
direction=1,两种情况
	1. 之前没有投过票，现在投赞成票   --> 更新分数和投票纪录
	2. 之前投反对票，现在改投赞成票   --> 更新分数和投票纪录
direction=0,两种情况
	1. 之前投过票，现在取消投赞成票   --> 更新分数和投票纪录
	2. 之前没投反对票，现在取消投票   --> 更新分数和投票纪录
direction=-1,两种情况
	1. 之前没有投过票，现在投反对票   --> 更新分数和投票纪录
	2. 之前投赞成票，现在改投反对票   --> 更新分数和投票纪录

投票限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了
	1. 到期之后将redis中保存的赞成票以及反对票存储到mysql表中
	2. 到期之后删除 KeyPostVotedZSetPF
*/

// VoteForPost 为帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
