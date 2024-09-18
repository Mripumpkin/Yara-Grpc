package main

import (
	"context"
	"net"

	"net/http"
	_ "net/http/pprof"
	"yara/common"
	grpcYARA "yara/grpclib"

	"google.golang.org/grpc"
)

type server_yara struct {
	grpcYARA.UnimplementedYARACheckServer
}

func (s *server_yara) YARACheck(ctx context.Context, in *grpcYARA.CheckRequest) (*grpcYARA.CheckReply, error) {
	common.Logger.Infof("YARA Check Sample Request:%v", in.Sample)
	str_location_res := common.YaraCheckOneSample(in.Sample, in.Ruleslist)
	return &grpcYARA.CheckReply{Reply: str_location_res}, nil
}

func (s *server_yara) YARAFileCheck(ctx context.Context, in *grpcYARA.FileRequest) (*grpcYARA.FileReply, error) {
	common.Logger.Infof("YARA_File Check Sample Request:%v", in.File)
	str_res := common.YaraFileCheckOneSample(in.File, in.Rules)
	return &grpcYARA.FileReply{Result: str_res}, nil
}

func main() {
	go func() {
		common.Logger.Info(http.ListenAndServe("localhost:6060", nil))
	}()
	// 监听本地端口
	_HOST := common.GetConfigFileHanlde().GetString("yara_grpc_server.host")
	_PORT := common.GetConfigFileHanlde().GetString("yara_grpc_server.port")
	listener, err := net.Listen("tcp", _HOST+":"+_PORT)
	if err != nil {
		common.Logger.Errorf("监听端口失败: %s", err)
		return
	}
	s := grpc.NewServer()
	grpcYARA.RegisterYARACheckServer(s, &server_yara{})
	common.Logger.Infof("yara_server 正在运行: %v", listener.Addr())
	err = s.Serve(listener)
	if err != nil {
		common.Logger.Errorf("开启服务失败: %s", err)
		return
	}
}
