package main

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
	bbPool "github.com/panjf2000/gnet/v2/pkg/pool/bytebuffer"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SayHello struct {
	*gnet.BuiltinEventEngine
}

// OnTraffic fires when a local socket receives data from the peer.
func (*SayHello) OnTraffic(c gnet.Conn) (action gnet.Action) {
	d, _ := c.Next(-1)
	fmt.Println("res ", string(d))
	return
}

func main() {
	var ops []gnet.Option
	ops = append(ops, func(opts *gnet.Options) {
		opts.TCPNoDelay = gnet.TCPNoDelay
	})

	client, _ := gnet.NewClient(&SayHello{}, ops...)

	conn, _ := client.Dial("tcp", fmt.Sprintf("%v:%v", "127.0.0.1", 8080))
	client.Start()

	i := 0
	for {
		buff := bbPool.Get()
		buff.SetString(fmt.Sprintf("hi %v", i))
		i = i + 1
		conn.Write(buff.Bytes())
		bbPool.Put(buff)
		time.Sleep(time.Millisecond * 100)
	}

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quitChan

}
