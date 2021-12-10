package main

import (
	"github.com/liangguifeng/shine/server"
	"shine-skeleton/startup"
)

const PORT = "8888"
// 此处定义当前服务的名字
const APP_NAME = "shine-skeleton"

func main() {
	s := server.New(
		server.WithHost("0.0.0.0"),
		server.WithPort(PORT),
		server.WithGRPCRegisterFunc(startup.RegisterGRPC),
		server.WithHTTPRegisterFunc(startup.RegisterHTTP))

	err := s.Start()
	if err != nil {
		return
	}
}
