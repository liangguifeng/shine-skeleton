package startup

import (
	"context"
	"github.com/liangguifeng/shine/server"
	"google.golang.org/grpc"
	hello_world "shine-skeleton/proto/shine_skeleton_proto/helloworld"
	"shine-skeleton/server/hello_world_server"
)

func RegisterHTTP(ctx context.Context, s *server.Server) {
	if err := hello_world.RegisterHelloWorldHandler(ctx, s.ServerMux, s.GRPClientConn); err != nil {
		panic(err)
	}
}

func RegisterGRPC(ctx context.Context, s *server.Server) {
	grpcServer := grpc.NewServer()
	hello_world.RegisterHelloWorldServer(grpcServer, hello_world_server.NewHelloWorldServer())
	if err := grpcServer.Serve(s.GRPCListener); err != nil {
		panic(err)
	}
}
