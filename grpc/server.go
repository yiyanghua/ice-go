package main

import (
	"io"
	"log"
	"net"
	"runtime"
	"strconv"

	"github.com/yiyanghua/ice/inf"
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

func (t *Data) Channel(stream inf.Data_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &inf.UserRp{
			Name: strconv.Itoa(int(args.Id)) + ":test",
		}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}
