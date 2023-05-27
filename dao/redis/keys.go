package redis

const (
	Prefix             = "liaoBa:"
	KeyPostTimeZSet    = "post:time"   // zset:帖子和发帖时间
	KeyPostScoreZSet   = "post:score"  // zset:帖子和投票分数
	KeyPostVotedZSetPF = "post:voted:" // zet:记录用户及投票类型;参数是post id
	KeyCommunitySetPF  = "community:"  // set; 保存每个分区下的帖子id
)

// 给redis key 加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
