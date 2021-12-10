package hello_world_server

import (
	"context"
	hello_world "shine-skeleton/proto/shine_skeleton_proto/helloworld"
)

type helloWorldServer struct {
}

func NewHelloWorldServer() *helloWorldServer {
	return &helloWorldServer{}
}

func (h helloWorldServer) HelloWorld(ctx context.Context, request *hello_world.HelloWorldRequest) (*hello_world.HelloWorldResponse, error) {
	return nil, nil
}
