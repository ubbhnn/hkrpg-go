package muip

import (
	"time"

	"github.com/gucooing/hkrpg-go/pkg/logger"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/protocol/cmd"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
	pb "google.golang.org/protobuf/proto"
)

const (
	PacketMaxLen = 343 * 1024 // 最大应用层包长度
)

func (s *Muip) NodeRegisterMessage(cmdId uint16, serviceMsg pb.Message) {
	switch cmdId {
	case cmd.ServiceConnectionRsp: // 注册回包
		s.ServiceConnectionRsp(serviceMsg)
	case cmd.GetServerOuterAddrRsp: // 心跳包
		s.GetServerOuterAddrRsp(serviceMsg)
	default:

	}
}

// 从node接收消息
func (s *Muip) RecvNode() {
	nodeMsg := make([]byte, PacketMaxLen)

	for {
		var bin []byte = nil
		recvLen, err := s.NodeConn.Read(nodeMsg)
		if err != nil {
			logger.Debug("exit recv loop, conn read err: %v", err)
			return
		}
		bin = nodeMsg[:recvLen]
		nodeMsgList := make([]*alg.PackMsg, 0)
		alg.DecodeBinToPayload(bin, &nodeMsgList, nil)
		for _, msg := range nodeMsgList {
			serviceMsg := alg.DecodePayloadToProto(msg)
			s.NodeRegisterMessage(msg.CmdId, serviceMsg)
		}
	}
}

// 发送到node
func (s *Muip) SendNode(cmdId uint16, playerMsg pb.Message) {
	rspMsg := new(alg.ProtoMsg)
	rspMsg.CmdId = cmdId
	rspMsg.PayloadMessage = playerMsg
	tcpMsg := alg.EncodeProtoToPayload(rspMsg)
	if tcpMsg.CmdId == 0 {
		logger.Error("cmdId error")
	}
	binMsg := alg.EncodePayloadToBin(tcpMsg, nil)
	_, err = s.NodeConn.Write(binMsg)
	if err != nil {
		logger.Debug("exit send loop, conn write err: %v", err)
		return
	}
}

// 向node注册
func (s *Muip) Connection() {
	req := &spb.ServiceConnectionReq{
		ServerType: spb.ServerType_SERVICE_MUIP,
		AppId:      s.AppId,
		Addr:       s.Config.OuterIp,
		Port:       s.Port,
	}

	s.SendNode(cmd.ServiceConnectionReq, req)
}

func (s *Muip) ServiceConnectionRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.ServiceConnectionRsp)
	if rsp.ServerType == spb.ServerType_SERVICE_MUIP && rsp.AppId == s.AppId {
		logger.Info("已向node注册成功！")
	}
	// 获取game地址/心跳包
	go s.GetServerOuterAddrReq()
}

func (s *Muip) GetServerOuterAddrReq() {
	// 心跳包
	for {
		req := &spb.GetServerOuterAddrReq{
			ServerType: spb.ServerType_SERVICE_MUIP,
			AppId:      s.AppId,
		}
		s.SendNode(cmd.GetServerOuterAddrReq, req)
		time.Sleep(time.Second * 5)
	}
}

func (s *Muip) GetServerOuterAddrRsp(serviceMsg pb.Message) {
	rsp := serviceMsg.(*spb.GetServerOuterAddrRsp)
	if rsp.ServerType != spb.ServerType_SERVICE_MUIP {
		return
	}
}
