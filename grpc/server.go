package main

import (
	"ice-go/inf"
	"log"
	"net"
	"runtime"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = "41005"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Println("grpc server in: %s", port)
}


// 定义方法
func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {
	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":test",
	}
	return response, err
}