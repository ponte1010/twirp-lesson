package main

import (
	"context"
	"net/http"

	pb "github.com/ponte1010/twirp-lesson/pkg/rpc/helloworld"
)

type HelloWorldServer struct{}

func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	return &pb.HelloResp{Text: "Hello " + req.GetSubject()}, nil
}

func main() {
	twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)
	http.ListenAndServe(":8080", mux)
}
