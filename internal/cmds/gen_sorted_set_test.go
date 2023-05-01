// Code generated DO NOT EDIT

package cmds

import "testing"

func sorted_set0(s Builder) {
	s.Bzmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Min().Count(1).Build()
	s.Bzmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Min().Build()
	s.Bzmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Max().Count(1).Build()
	s.Bzmpop().Timeout(1).Numkeys(1).Key("1").Key("1").Max().Build()
	s.Bzpopmax().Key("1").Key("1").Timeout(1).Build()
	s.Bzpopmin().Key("1").Key("1").Timeout(1).Build()
	s.Zadd().Key("1").Nx().Gt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Gt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Gt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Gt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Lt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Lt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Lt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Lt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Nx().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Gt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Gt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Gt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Gt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Lt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Lt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Lt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Lt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Xx().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Gt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Gt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Gt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Gt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Lt().Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Lt().Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Lt().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Lt().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Ch().Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Ch().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").Incr().ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zadd().Key("1").ScoreMember().ScoreMember(1, "1").ScoreMember(1, "1").Build()
	s.Zcard().Key("1").Build()
	s.Zcard().Key("1").Cache()
	s.Zcount().Key("1").Min("1").Max("1").Build()
	s.Zcount().Key("1").Min("1").Max("1").Cache()
	s.Zdiff().Numkeys(1).Key("1").Key("1").Withscores().Build()
	s.Zdiff().Numkeys(1).Key("1").Key("1").Build()
	s.Zdiffstore().Destination("1").Numkeys(1).Key("1").Key("1").Build()
	s.Zincrby().Key("1").Increment(1).Member("1").Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateSum().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateSum().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateMin().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateMin().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateMax().Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").AggregateMax().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Withscores().Build()
	s.Zinter().Numkeys(1).Key("1").Key("1").Build()
	s.Zintercard().Numkeys(1).Key("1").Key("1").Limit(1).Build()
	s.Zintercard().Numkeys(1).Key("1").Key("1").Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateSum().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateMin().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateMax().Build()
	s.Zinterstore().Destination("1").Numkeys(1).Key("1").Key("1").Build()
	s.Zlexcount().Key("1").Min("1").Max("1").Build()
	s.Zlexcount().Key("1").Min("1").Max("1").Cache()
	s.Zmpop().Numkeys(1).Key("1").Key("1").Min().Count(1).Build()
	s.Zmpop().Numkeys(1).Key("1").Key("1").Min().Build()
	s.Zmpop().Numkeys(1).Key("1").Key("1").Max().Count(1).Build()
	s.Zmpop().Numkeys(1).Key("1").Key("1").Max().Build()
	s.Zmscore().Key("1").Member("1").Member("1").Build()
	s.Zmscore().Key("1").Member("1").Member("1").Cache()
	s.Zpopmax().Key("1").Count(1).Build()
	s.Zpopmax().Key("1").Build()
	s.Zpopmin().Key("1").Count(1).Build()
	s.Zpopmin().Key("1").Build()
	s.Zrandmember().Key("1").Count(1).Withscores().Build()
	s.Zrandmember().Key("1").Count(1).Build()
	s.Zrandmember().Key("1").Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Limit(1, 1).Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Rev().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Limit(1, 1).Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Build()
	s.Zrange().Key("1").Min("1").Max("1").Byscore().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Limit(1, 1).Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Rev().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Limit(1, 1).Withscores().Cache()
}

func sorted_set1(s Builder) {
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Build()
	s.Zrange().Key("1").Min("1").Max("1").Bylex().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Limit(1, 1).Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Build()
	s.Zrange().Key("1").Min("1").Max("1").Rev().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Limit(1, 1).Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Limit(1, 1).Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Limit(1, 1).Build()
	s.Zrange().Key("1").Min("1").Max("1").Limit(1, 1).Cache()
	s.Zrange().Key("1").Min("1").Max("1").Withscores().Build()
	s.Zrange().Key("1").Min("1").Max("1").Withscores().Cache()
	s.Zrange().Key("1").Min("1").Max("1").Build()
	s.Zrange().Key("1").Min("1").Max("1").Cache()
	s.Zrangebylex().Key("1").Min("1").Max("1").Limit(1, 1).Build()
	s.Zrangebylex().Key("1").Min("1").Max("1").Limit(1, 1).Cache()
	s.Zrangebylex().Key("1").Min("1").Max("1").Build()
	s.Zrangebylex().Key("1").Min("1").Max("1").Cache()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Withscores().Limit(1, 1).Build()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Withscores().Limit(1, 1).Cache()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Withscores().Build()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Withscores().Cache()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Limit(1, 1).Build()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Limit(1, 1).Cache()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Build()
	s.Zrangebyscore().Key("1").Min("1").Max("1").Cache()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Byscore().Rev().Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Byscore().Rev().Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Byscore().Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Byscore().Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Bylex().Rev().Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Bylex().Rev().Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Bylex().Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Bylex().Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Rev().Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Rev().Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Limit(1, 1).Build()
	s.Zrangestore().Dst("1").Src("1").Min("1").Max("1").Build()
	s.Zrank().Key("1").Member("1").Withscore().Build()
	s.Zrank().Key("1").Member("1").Withscore().Cache()
	s.Zrank().Key("1").Member("1").Build()
	s.Zrank().Key("1").Member("1").Cache()
	s.Zrem().Key("1").Member("1").Member("1").Build()
	s.Zremrangebylex().Key("1").Min("1").Max("1").Build()
	s.Zremrangebyrank().Key("1").Start(1).Stop(1).Build()
	s.Zremrangebyscore().Key("1").Min("1").Max("1").Build()
	s.Zrevrange().Key("1").Start(1).Stop(1).Withscores().Build()
	s.Zrevrange().Key("1").Start(1).Stop(1).Withscores().Cache()
	s.Zrevrange().Key("1").Start(1).Stop(1).Build()
	s.Zrevrange().Key("1").Start(1).Stop(1).Cache()
	s.Zrevrangebylex().Key("1").Max("1").Min("1").Limit(1, 1).Build()
	s.Zrevrangebylex().Key("1").Max("1").Min("1").Limit(1, 1).Cache()
	s.Zrevrangebylex().Key("1").Max("1").Min("1").Build()
	s.Zrevrangebylex().Key("1").Max("1").Min("1").Cache()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Withscores().Limit(1, 1).Build()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Withscores().Limit(1, 1).Cache()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Withscores().Build()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Withscores().Cache()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Limit(1, 1).Build()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Limit(1, 1).Cache()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Build()
	s.Zrevrangebyscore().Key("1").Max("1").Min("1").Cache()
	s.Zrevrank().Key("1").Member("1").Withscore().Build()
	s.Zrevrank().Key("1").Member("1").Withscore().Cache()
	s.Zrevrank().Key("1").Member("1").Build()
	s.Zrevrank().Key("1").Member("1").Cache()
	s.Zscan().Key("1").Cursor(1).Match("1").Count(1).Build()
	s.Zscan().Key("1").Cursor(1).Match("1").Build()
	s.Zscan().Key("1").Cursor(1).Count(1).Build()
	s.Zscan().Key("1").Cursor(1).Build()
	s.Zscore().Key("1").Member("1").Build()
	s.Zscore().Key("1").Member("1").Cache()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateSum().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateSum().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateMin().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateMin().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateMax().Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").AggregateMax().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Withscores().Build()
	s.Zunion().Numkeys(1).Key("1").Key("1").Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateSum().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMin().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).AggregateMax().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").Weights(1).Weights(1).Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateSum().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateMin().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").AggregateMax().Build()
	s.Zunionstore().Destination("1").Numkeys(1).Key("1").Key("1").Build()
}

func TestCommand_InitSlot_sorted_set(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { sorted_set0(s) })
	t.Run("1", func(t *testing.T) { sorted_set1(s) })
}

func TestCommand_NoSlot_sorted_set(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { sorted_set0(s) })
	t.Run("1", func(t *testing.T) { sorted_set1(s) })
}