package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"ice-go/grpc/inf"
)

var (
	wg sync.WaitGroup
)

const (
	server   = "127.0.0.1"
	parallel = 5 //连接并行度
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	/*	//并行请求
		for i := 0; i < parallel; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				exe()
			}()
		}
		wg.Wait()*/

	stream()

	log.Printf("time taken: %.2f ", time.Now().Sub(currTime).Seconds())
}

func exe() {
	//建立连接
	conn, _ := grpc.Dial(server+":41005", grpc.WithInsecure())
	defer conn.Close()
	client := inf.NewDataClient(conn)
	getUser(client)
}

func stream() {
	//建立连接
	conn, _ := grpc.Dial(server+":41005", grpc.WithInsecure())
	defer conn.Close()
	client := inf.NewDataClient(conn)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			var request inf.UserRq
			r := rand.Intn(parallel)
			request.Id = int32(r)

			if err := stream.Send(&request); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply.Name)
	}
}

func getUser(client inf.DataClient) {
	var request inf.UserRq
	r := rand.Intn(parallel)
	request.Id = int32(r)

	response, _ := client.GetUser(context.Background(), &request) //调用远程方法

	//判断返回结果是否正确
	if id, _ := strconv.Atoi(strings.Split(response.Name, ":")[0]); id != r {
		log.Printf("response error  %#v", response)
	} else {
		log.Printf("response %#v", response)
	}

}
