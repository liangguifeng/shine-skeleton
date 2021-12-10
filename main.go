package main

import (
	"github.com/liangguifeng/shine/server"
	"shine-skeleton/startup"
)

const PORT = ":8888"

func main() {
	s := server.New(
		server.WithEndpoint(PORT),
		server.WithGRPCRegisterFunc(startup.RegisterGRPC),
		server.WithHTTPRegisterFunc(startup.RegisterHTTP))

	err := s.Start()
	if err != nil {
		return
	}
}
