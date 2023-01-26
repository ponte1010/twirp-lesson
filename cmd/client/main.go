package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "github.com/ponte1010/twirp-lesson/pkg/rpc/helloworld"
)

func main() {
	client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

	subject := ""
	if len(os.Args) > 1 {
		subject = os.Args[1]
	}
	resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: subject})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetText())
}
