package gate

import (
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gucooing/hkrpg-go/gateserver/config"
	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/random"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

var syncGD sync.Mutex

func (s *GateServer) HandlePlayerGetTokenCsReq(p *PlayerGame, playerMsg []byte) {
	req := new(proto.PlayerGetTokenCsReq)
	pb.Unmarshal(playerMsg, req)
	rsp := new(proto.PlayerGetTokenScRsp)

	// 人数验证
	if config.GetConfig().MaxPlayer != -1 {
		if CLIENT_CONN_NUM >= config.GetConfig().MaxPlayer {
			rsp.Uid = 0
			rsp.Retcode = uint32(proto.Retcode_RET_REACH_MAX_PLAYER_NUM)
			rsp.Msg = "当前服务器人数过多，请稍后再试。"
			GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
			return
		}
	}

	// 请求验证
	if req.Token == "" || req.AccountUid == "" {
		return
	}
	accountUid, err := strconv.ParseUint(req.AccountUid, 10, 64)
	if err != nil {
		logger.Error("get token uid error")
		return
	}
	uidPlayer := s.Store.QueryUidPlayerUidByFieldPlayer(uint32(accountUid))

	// token验证
	if uidPlayer.ComboToken != req.Token {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_ACCOUNT_VERIFY_ERROR)
		rsp.Msg = "token验证失败"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,token验证失败", accountUid)
		return
	}

	// 封禁验证
	if uidPlayer.IsBan {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_IN_GM_BIND_ACCESS)
		rsp.Msg = "该账号正处于封禁状态，暂时无法登录，详情可联系客服。"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Info("登录账号:%v,已被封禁", accountUid)
		return
	}

	p.Uid = uint32(uidPlayer.AccountUid)

	newuidPlayer := &UidPlayer{
		AccountId:  uidPlayer.AccountId,
		IsBan:      false,
		ComboToken: "",
	}

	syncGD.Lock()
	err = s.Store.UpdateUidPlayer(uidPlayer.AccountId, newuidPlayer)
	syncGD.Unlock()
	if err != nil {
		rsp.Uid = 0
		rsp.Retcode = uint32(proto.Retcode_RET_PLAYER_DATA_ERROR)
		rsp.Msg = "玩家数据错误"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("登录账号:%v,玩家数据错误", accountUid)
		return
	}

	p.IsToken = true

	// 登录成功，拉取game
	if s.gameAppId == "" || s.gameAll[s.gameAppId] == nil {
		rsp.Uid = p.Uid
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		return
	}

	gameAppId := s.GetGameAppId()
	game := s.gameAll[gameAppId]
	if game == nil {
		rsp.Uid = p.Uid
		rsp.Retcode = uint32(proto.Retcode_RET_SYSTEM_BUSY)
		rsp.Msg = "game未启动"
		GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
		logger.Error("game未启动")
		return
	}

	p.NewGame(game.addr)
	p.GameAppId = game.appId
	gamereq := &spb.PlayerLoginReq{
		PlayerUid: p.Uid,
		AppId:     s.gameAppId,
	}
	p.sendGame(cmd.PlayerLoginReq, gamereq)

	// 本gate重复登录验证//不向node发送玩家登录通知
	if player, ok := GAMESERVER.sessionMap[p.Uid]; ok {
		notify := new(proto.GetChallengeScRsp)
		// TODO 是的，没错，还是同样的原因
		// 重复登录下线通知
		GateToPlayer(player, cmd.PlayerKickOutScNotify, notify)
		// 删除连接
		player.GameConn.Close()
	} else {
		// 异步通知给node
		go s.sendNode(cmd.PlayerLoginReq, gamereq)
	}

	// 构造回复内容
	timeRand := random.GetTimeRand()
	serverSeedUint64 := timeRand.Uint64()
	p.Seed = serverSeedUint64
	rsp.Uid = p.Uid
	rsp.SecretKeySeed = serverSeedUint64
	rsp.BlackInfo = &proto.BlackInfo{}
	GateToPlayer(p, cmd.PlayerGetTokenScRsp, rsp)
}

func (p *PlayerGame) NewGame(gameAddr string) {
	gameConn, err := net.Dial("tcp", gameAddr)
	if err != nil {
		logger.Error("无法连接到GAME:", err)
		return
	}
	p.GameConn = gameConn
	go p.recvGame()
}

func (s *GateServer) GetGameAppId() string {
	gameAppId := s.gameAppId

	for _, appId := range s.errGameAppId {
		if gameAppId == appId {
			gameAppId = s.GetMinGameAppId(appId)
		}
	}

	return gameAppId
}

func (s *GateServer) GetMinGameAppId(errAppId string) string {
	var minNum uint64
	var minAppId string
	for _, game := range s.gameAll {
		if game.appId == errAppId {
			continue
		}
		if minAppId == "" || minNum > game.num {
			minAppId = game.appId
			minNum = game.num
		}
	}
	return minAppId
}

func (p *PlayerGame) HandlePlayerHeartBeatCsReq(payloadMsg []byte) {
	req := new(proto.PlayerHeartbeatCsReq)
	pb.Unmarshal(payloadMsg, req)

	rsp := new(proto.PlayerHeartbeatScRsp)
	rsp.ServerTimeMs = uint64(time.Now().UnixNano() / 1e6)
	rsp.ClientTimeMs = req.ClientTimeMs

	GateToPlayer(p, cmd.PlayerHeartBeatScRsp, rsp)
}
