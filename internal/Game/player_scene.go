package Game

import (
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
)

func (g *Game) GetRogueScoreRewardInfoCsReq() {
	rsp := new(proto.GetRogueScoreRewardInfoScRsp)
	rsp.ScoreRewardInfo = &proto.RogueScoreRewardInfo{
		HasTakenInitialScore: true,
		PoolRefreshed:        true,
		PoolId:               20,
	}

	g.send(cmd.GetRogueScoreRewardInfoScRsp, rsp)
}
